package main

import (
	"Go_Learn/RpcDemo/codec"
	"sync"
)

type Client struct {
	clientCodec codec.ClientCodec
	sending     sync.Mutex
	pending     map[uint64]string
}

func (c *Client) registerCall() {

}

func (c *Client) removeCall() {

}

func (c *Client) receive() {

}

func (c *Client) send() {

}

func (c *Client) Close() {

}

func Dial() {

}
