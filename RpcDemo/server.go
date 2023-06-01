package RpcDemo

import (
	"Go_Learn/RpcDemo/codec"
	"sync"
)

type Server struct {
	serverCodec codec.ServerCodec
	sending     sync.Mutex
	pending     map[string]*service
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) serveCodec() {

}

func (s *Server) ServeConn() {

}

func (s *Server) handleRequest() {

}

func (s *Server) Register() {

}

func (s *Server) Accept() {

}
