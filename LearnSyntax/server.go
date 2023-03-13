package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 所有以/开头的请求都会路由到handler函数
	http.HandleFunc("/", handler)
	// 监听localhost:8000端口如果有错误就打印错误并退出
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// 该函数负责将发起请求的报文相关信息打印出来
func handler(w http.ResponseWriter, r *http.Request) {
	//%q用于将给定字符串转义并以双引号包裹输出具体来说，%q会将字符串中的特殊字符进行转义，
	//例如换行符、制表符和双引号等，并添加双引号包裹整个字符串
	fmt.Fprintf(w, "Hello World！This is response\n")
	fmt.Fprintf(w, "URL.Path is = %q\n", r.URL.Path)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
