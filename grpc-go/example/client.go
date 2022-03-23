/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/3/14 9:45 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"math/rand"
	"time"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func printFeature(client RouteGuideClient, point *Point) {
	log.Printf("Getting feature for point (%d, %d)", point.Latitude, point.Longitude)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	feature, err := client.GetFeature(ctx, point)
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
	}
	log.Println(feature)
}


func printFeatures(client RouteGuideClient, rect *Rectangle) {
	log.Printf("Looking for features within %v", rect)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.ListFeatures(ctx, rect)
	if err != nil {
		log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Printf("Feature: name: %q, point:(%v, %v)", feature.GetName(),
			feature.GetLocation().GetLatitude(), feature.GetLocation().GetLongitude())
	}
}

func runRecordRoute(client RouteGuideClient) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pointCount := int(r.Int31n(100)) + 2 // Traverse at least two points
	var points []*Point
	for i := 0; i < pointCount; i++ {
		points = append(points, randomPoint(r))
	}
	log.Printf("Traversing %d points.", len(points))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.RecordRoute(ctx)
	if err != nil {
		log.Fatalf("%v.RecordRoute(_) = _, %v", client, err)
	}
	for _, point := range points {
		if err := stream.Send(point); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, point, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Route summary: %v", reply)
}


func runRouteChat(client RouteGuideClient) {
	notes := []*RouteNote{
		{Location: &Point{Latitude: 0, Longitude: 1}, Message: "First message"},
		{Location: &Point{Latitude: 0, Longitude: 2}, Message: "Second message"},
		{Location: &Point{Latitude: 0, Longitude: 3}, Message: "Third message"},
		{Location: &Point{Latitude: 0, Longitude: 1}, Message: "Fourth message"},
		{Location: &Point{Latitude: 0, Longitude: 2}, Message: "Fifth message"},
		{Location: &Point{Latitude: 0, Longitude: 3}, Message: "Sixth message"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.RouteChat(ctx)
	if err != nil {
		log.Fatalf("%v.RouteChat(_) = _, %v", client, err)
	}
	waitc := make(chan struct{})

	for _, note := range notes {
		if err := stream.Send(note); err != nil {
			log.Fatalf("Failed to send a note: %v", err)
		}
	}
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Got message %s at point(%d, %d)", in.Message, in.Location.Latitude, in.Location.Longitude)
		}
	}()
	_ = stream.CloseSend()
	<-waitc
}

func randomPoint(r *rand.Rand) *Point {
	lat := (r.Int31n(180) - 90) * 1e7
	long := (r.Int31n(360) - 180) * 1e7
	return &Point{Latitude: lat, Longitude: long}
}


func main() {
	var opts []grpc.DialOption
	if *tls {
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer func() {
		_ = conn.Close()
	}()
	client := NewRouteGuideClient(conn)


	// Looking for a valid feature
	printFeature(client, &Point{Latitude: 409146138, Longitude: -746188906})

	// Feature missing.
	printFeature(client, &Point{Latitude: 0, Longitude: 0})

	// Looking for features between 40, -75 and 42, -73.
	printFeatures(client, &Rectangle{
		Lo: &Point{Latitude: 400000000, Longitude: -750000000},
		Hi: &Point{Latitude: 420000000, Longitude: -730000000},
	})

	// RecordRoute
	runRecordRoute(client)

	// RouteChat
	runRouteChat(client)

}
