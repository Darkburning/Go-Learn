package codec

import (
	"Go_Learn/RpcDemo/serializer"
	"bufio"
	"io"
	"net"
)

// ClientCodec 持有连接
type ClientCodec struct {
	conn       io.ReadWriteCloser
	serializer serializer.JsonSerializer
	w          *bufio.Writer
	r          *bufio.Reader
}

func NewClientCodec(conn net.Conn) *ClientCodec {
	return &ClientCodec{
		conn: conn,
		w:    bufio.NewWriter(conn),
		r:    bufio.NewReader(conn),
	}
}

// 实现ReadRespHeader、ReadRespBody、WriteReq、Close

func (c *ClientCodec) ReadResponse() {

}
func (c *ClientCodec) WriteRequest() {
	defer func(w *bufio.Writer) {
		err := w.Flush()
		if err != nil {

		}
	}(c.w)
	msg := "Hello world!"
	byteMsg, err := c.serializer.Marshal(msg)
	if err != nil {

	}
	err = sendFrame(c.w, byteMsg)
	if err != nil {
		return
	}

}
func (c *ClientCodec) Close() {

}
