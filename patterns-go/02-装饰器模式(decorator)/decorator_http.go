/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/17 3:23 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _2_装饰器模式_decorator_

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)


// 写HTTP响应头的修饰器，其传入一个 http.HandlerFunc，然后返回一个改写的版本
func WithServerHeader(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("--->WithServerHeader()")
		w.Header().Set("Server", "HelloServer v0.0.1")
		h(w, r)
	}
}

// 写认证 Cookie 的修饰器
func WithAuthCookie(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("--->WithAuthCookie()")
		cookie := &http.Cookie{Name: "Auth", Value: "Pass", Path: "/"}
		http.SetCookie(w, cookie)
		h(w, r)
	}
}

// 检查认证Cookie的 修饰器
func WithBasicAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("--->WithBasicAuth()")
		cookie, err := r.Cookie("Auth")
		if err != nil || cookie.Value != "Pass" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		h(w, r)
	}
}

// 打日志的修饰器
func WithDebugLog(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("--->WithDebugLog")
		_ = r.ParseForm()
		log.Println(r.Form)
		log.Println("path", r.URL.Path)
		log.Println("scheme", r.URL.Scheme)
		log.Println(r.Form["url_long"])
		for k, v := range r.Form {
			log.Println("key:", k)
			log.Println("val:", strings.Join(v, ""))
		}
		h(w, r)
	}
}


func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Recieved Request %s from %s\n", r.URL.Path, r.RemoteAddr)
	_, _ = fmt.Fprintf(w, "Hello, World! "+r.URL.Path)
}