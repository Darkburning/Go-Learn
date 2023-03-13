package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// 在基础功能上新增计数功能
// 如果请求的pattern是以/结尾，那么所有以该url为前缀的url都会被这条规则匹配。
// 在这些代码的背后，服务器每一次接收请求处理时都会另起一个goroutine，
// 这样服务器就可以同一时间处理多个请求。然而在并发情况下，假如真的有两个请求同一时刻去更新count，那么这个值可能并不会被正确地增加；
// 这个程序可能会引发一个严重的bug：竞态条件
// var mu sync.Mutex 通过使用sync.Mutex实现互斥锁，用于保护共享变量count的线程安全。使用mu.Lock()和mu.Unlock()来分别锁定和解锁锁定count变量的访问
// 声明全局变量
var mu sync.Mutex
var count int

func main() {
	// 所有以/开头的pattern都会路由到normalHandler包括/count
	http.HandleFunc("/", normalHandler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func normalHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "\033[33mHello World!\nURL.Path = %q\033[0m\n", r.URL.Path) //ANSI escape码控制终端文本的显示
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "\033[31mcount is:%v\033[0m\n", count) //ANSI escape码控制终端文本的显示
	mu.Unlock()
}
