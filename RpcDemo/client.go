package main

import (
	"Go_Learn/RpcDemo/codec"
	"Go_Learn/RpcDemo/protocol"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

type Client struct {
	clientCodec *codec.ClientCodec
	sending     *sync.Mutex
}

func (c *Client) Close() error {
	return c.clientCodec.Close()
}

func (c *Client) Call(method string, args ...interface{}) []interface{} {
	c.sending.Lock()

	req := &protocol.Request{
		Method: method,
		Args:   args,
	}
	c.clientCodec.WriteRequest(req)
	c.sending.Unlock()

	var err error
	for err == nil {
		resp, err := c.clientCodec.ReadResponse()
		if err != nil {
			log.Println("rpc client: client receive: " + err.Error())
		}
		if resp.Err != "" {
			log.Println("rpc client: client receive: " + err.Error())
			return nil
		} else {
			fmt.Printf("client call success!\n")
			for idx, reply := range resp.Replies {
				fmt.Printf("Value %d is : %v\n", idx, reply)
			}
			return resp.Replies
		}
	}
	return nil
}

// Dial 处理建立连接超时
func Dial(network string, addr string) (*Client, error) {
	conn, err := net.DialTimeout(network, addr, timeOutLimit)
	if err != nil {
		return nil, err
	} else {
		defer func() {
			if err != nil {
				_ = conn.Close()
			}
		}()
		// 创建子协程，创建一个客户端
		ch := make(chan *Client)
		go func() {
			ch <- &Client{
				clientCodec: codec.NewClientCodec(conn),
				sending:     new(sync.Mutex),
			}
		}()

		// select多路复用处理阻塞IO，在两个信道上监听
		select {
		case <-time.After(timeOutLimit): // New Client Timeout
			return nil, errors.New("rpc client: dial timeout: expect within 5s")
		case result := <-ch: // New Client Success
			return result, nil
		}

	}
}
