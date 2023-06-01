package rpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func startServer(addr chan string) {
	err := rpc.RegisterName("HelloService", new(HelloService))
	if err != nil {
		return
	}

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	addr <- listener.Addr().String()

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}

func startMain() {
	addr := make(chan string)
	go startServer(addr)
	address := <-addr

	client, err := rpc.Dial("tcp", address)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "world", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
