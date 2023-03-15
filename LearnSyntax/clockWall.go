package main

import (
	"io"
	"log"
	"net"
	"os"
)

// 向不同时区的两个服务器发起请求并打印响应结果
func main() {
	Dial(1)
	Dial(2)
}
func Dial(i int) {
	portNumber := os.Args[i]
	addr := "localhost:" + portNumber
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// 将tcp通信获得的信息写入标准输出流
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
