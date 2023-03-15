package main

import (
	"io"
	"log"
	"net"
	"time"
)

// 创建一个返回伦敦时间的服务器
func main() {
	Listener, err := net.Listen("tcp", "localhost:8010")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := Listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		// 先获得对应地区再得出时间
		londonTimeZone, _ := time.LoadLocation("Europe/London")
		now := time.Now()
		londonTime := now.In(londonTimeZone)
		// 往连接里写入London时间
		_, err := io.WriteString(conn, "LondonTime is:"+londonTime.Format("15:04:05\n"))
		if err != nil {
			log.Print(err)
			return
		}
		time.Sleep(time.Second)
	}

}
