/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/28 10:55 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package main

import (
	"fmt"
	"time"
)

func logs(args ...interface{}) {
	fmt.Println(args...)
}

func tokenBucket(limit int, rate int) chan struct{} {
	tb := make(chan struct{}, limit)
	ticker := time.NewTicker(time.Duration(1) * time.Second)
	for i := 0; i < limit; i++ {
		tb <- struct{}{}
	}

	go func() {
		for {
			for i := 0; i < rate; i++ {
				tb <- struct{}{}
			}
			<-ticker.C
		}
	}()
	return tb
}

func popToken(bucket chan struct{}, n int) {
	for i := 0; i < n; i++ {
		<-bucket
	}
}

func testTokenBucket() {
	rate := 10
	tb := tokenBucket(20, rate)

	dataLen := 100
	for i := 0; i <= dataLen; i += rate {
		popToken(tb, rate)
		logs(i)
	}
}

func main() {
	testTokenBucket()
}
