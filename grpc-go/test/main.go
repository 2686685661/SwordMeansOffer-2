/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/3/14 8:09 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package main

import (
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	test := &Student{
		Name:                 "geetkutu",
		Male:                 true,
		Scores:               []int32{29,292,69},
	}
	data, err := proto.Marshal(test)
	if nil != err {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &Student{}
	err = proto.Unmarshal(data, newTest)
	if nil != err {
		log.Fatal("unmarshaling error: ", err)
	}
	if test.GetName() != newTest.GetName() {
		log.Fatalf("data mismatch %q != %q", test.GetName(), newTest.GetName())
	}

}


