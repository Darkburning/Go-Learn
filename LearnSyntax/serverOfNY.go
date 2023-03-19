package main

import (
	"io"
	"log"
	"net"
	"time"
)

// 创建一个返回纽约时间的服务器
func main() {
	Listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := Listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

}
func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		// 先获得对应地区再得出时间
		newYorkTimeZone, _ := time.LoadLocation("America/New_York")
		newYorkTime := time.Now().In(newYorkTimeZone)
		_, err := io.WriteString(conn, "NewYorkTime is:"+newYorkTime.Format("15:04:05\n"))
		if err != nil {
			log.Print(err)
			return
		}
		time.Sleep(time.Second)
	}
}
