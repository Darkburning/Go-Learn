package main

import (
	"Go_Learn/RpcDemo/codec"
	"errors"
	"io"
	"log"
	"net"
	"reflect"
	"sync"
)

type Server struct {
	sending *sync.Mutex              // 保证线程安全
	pending map[string]reflect.Value // 维护的service列表
}

func NewServer() *Server {
	return &Server{
		sending: new(sync.Mutex),
		pending: make(map[string]reflect.Value),
	}
}

// serveCodec 流程：读取请求，处理请求，发送响应
func (s *Server) serveCodec(sc *codec.ServerCodec) {
	s.sending.Lock()
	defer s.sending.Unlock()
	wg := new(sync.WaitGroup) // 等待直到所有的请求被处理完
	for {
		req, err := sc.ReadRequest()
		if err != nil {
			if req == nil {
				break // 尽力而为只有在header解析失败才终止循环
			}
			// 发送读取错误的报文
			sc.WriteResponse(err, errors.New("error When ReadRequest"))
			continue
		}
		wg.Add(1) // 需等待的协程+1
		//go s.handleRequest(sc, req, s.sending, wg) // 利用协程并发处理请求
	}
	wg.Wait() // 等待所有请求的处理结束
	_ = sc.Close()
}

func (s *Server) ServeConn(conn io.ReadWriteCloser) {
	defer func() {
		err := conn.Close()
		if err != nil {
			return
		}
	}()
	// 每个连接一个，负责编解码并读取数据
	serverCodec := codec.NewServerCodec(conn)

	s.serveCodec(serverCodec)
}

//func (s *Server) handleRequest(sc *codec.ServerCodec, req *protocol.Request, wg *sync.WaitGroup) {
//	defer wg.Done()
//	called := make(chan struct{})
//	sent := make(chan struct{})
//
//	go func() {
//		svc := s.pending[req.Method]
//		// 根据req获取入参列表
//
//		svc.Call(inArgs)
//		called <- struct{}{}
//		if err != nil {
//			req.h.Error = err.Error()
//			s.WriteResponse()
//			sent <- struct{}{}
//			return
//		}
//		// 发送响应报文
//		sc.WriteResponse(cc, req.h, req.replyv.Interface(), sending)
//		sent <- struct{}{}
//	}()
//
//	select {
//	case <-time.After(time.Second * 5):
//		req.h.Error = fmt.Sprintf("rpc server: request handle timeout: expect within 5s")
//		s.WriteResponse(cc, req.h, invalidRequest, sending)
//	case <-called:
//		<-sent
//	}
//}

// Register 传入结构体将其内所有方法注册到pending[string]reflect.Value
func (s *Server) Register(serviceName string, f interface{}) {
	if _, ok := s.pending[serviceName]; ok {
		log.Println("Already Registered!")
		return
	}

	fVal := reflect.ValueOf(f)
	s.pending[serviceName] = fVal

}

func (s *Server) isMethodExists(method string) bool {
	if _, ok := s.pending[method]; ok {
		return true
	} else {
		return false
	}
}

func (s *Server) call(serviceName string, inArgs []reflect.Value) []reflect.Value {
	if !s.isMethodExists(serviceName) {
		return nil
	} else {
		return s.pending[serviceName].Call(inArgs)
	}
}

// Accept 方法实现接收监听者的连接 开启协程处理每个到来的连接
// 若想启动服务只需传入listener，TCP或UNIX协议均可
func (s *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("rpc server: accept error:", err)
		}
		go s.ServeConn(conn)
	}
}
