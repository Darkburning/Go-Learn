package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// 当一个goroutine尝试在一个channel上做send或者receive操作时，这个goroutine会阻塞在调用处，
// 直到另一个goroutine从这个channel里接收或者写入值，这样两个goroutine才会继续执行channel操作之后的逻辑
// 在这个例子中，每一个fetch函数在执行时都会往channel里发送一个值(ch <- expression)，主函数负责接收这些值(<-ch)。
// 这个程序中我们用main函数来接收所有fetch函数传回的字符串，可以避免在goroutine异步执行还没有完成时main函数提前退出
func main() {
	start := time.Now()
	ch := make(chan string)
	// 通过for循环开启协程
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // go function()开启一个goroutine
	}
	// 依次打印出channel获得的字符串
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	// 整个程序执行的时间
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// 该函数接受url并通过一个string类型的chan传递获取到的url的报文主体的长度及耗时
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil { //有错误将错误传入通道并返回
		ch <- fmt.Sprint(err)
		return
	}

	//无错误则获得响应报文的body部分的长度
	lenOfBody, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil { //获取响应报文的主体出错，将错误的url及错误传回通道中并返回
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	// 如果一切无误则将耗时及报文长度及URL传入通道内
	ch <- fmt.Sprintf("%.3fs %7d %s", secs, lenOfBody, url)
}
