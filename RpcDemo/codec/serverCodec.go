package codec

import (
	"Go_Learn/RpcDemo/serializer"
	"bufio"
	"io"
	"net"
)

type ServerCodec struct {
	conn       io.ReadWriteCloser
	serializer serializer.JsonSerializer
	w          *bufio.Writer
	r          *bufio.Reader
}

func NewServerCodec(conn net.Conn) *ServerCodec {
	return &ServerCodec{
		conn: conn,
		w:    bufio.NewWriter(conn),
		r:    bufio.NewReader(conn),
	}
}

// 实现ReadReqHeader、ReadReqBody、WriteResp、Close

func (c *ServerCodec) ReadRequest() {

}

func (c *ServerCodec) WriteResponse() {

}

func (c *ServerCodec) Close() {

}
