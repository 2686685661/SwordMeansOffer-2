/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/17 3:25 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _2_装饰器模式_decorator_

import (
	"log"
	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {
	http.HandleFunc("/v1/hello", WithServerHeader(WithAuthCookie(hello)))
	http.HandleFunc("/v2/hello", WithServerHeader(WithBasicAuth(hello)))
	http.HandleFunc("/v3/hello", WithServerHeader(WithBasicAuth(WithDebugLog(hello))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
