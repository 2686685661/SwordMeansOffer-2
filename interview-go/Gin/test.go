/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/12 12:04 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Index"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}