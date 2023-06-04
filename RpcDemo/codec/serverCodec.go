package codec

import (
	"Go_Learn/RpcDemo/protocol"
	"Go_Learn/RpcDemo/serializer"
	"bufio"
	"io"
	"log"
)

type ServerCodec struct {
	conn       io.ReadWriteCloser
	serializer serializer.JsonSerializer
	w          *bufio.Writer
	r          *bufio.Reader
}

func NewServerCodec(conn io.ReadWriteCloser) *ServerCodec {
	return &ServerCodec{
		conn: conn,
		w:    bufio.NewWriter(conn),
		r:    bufio.NewReader(conn),
	}
}

// 实现ReadReqHeader、ReadReqBody、WriteResp、Close

func (c *ServerCodec) ReadRequest() (*protocol.Request, error) {
	var req protocol.Request
	byteReq, err := recvFrame(c.r)
	if err != nil {
		log.Fatal("ReadRequest Failed!")
		return nil, err
	}
	err = c.serializer.Unmarshal(byteReq, req)
	if err != nil {
		log.Fatal(err)
	}
	return &req, nil
}

func (c *ServerCodec) WriteResponse(errMsg error, replies interface{}) {
	defer func() {
		err := c.w.Flush() // 将所有的缓存数据写入底层的IO接口
		if err != nil {
			_ = c.Close() // 发生错误则关闭
		}
	}()
	errMsgBytes, err := c.serializer.Marshal(errMsg)
	if err != nil {
		log.Fatal(err)
	}
	repliesBytes, e := c.serializer.Marshal(replies)
	if e != nil {
		log.Fatal(e)
	}
	data := append(errMsgBytes, repliesBytes...)

	err = sendFrame(c.w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *ServerCodec) Close() error {
	return c.conn.Close()
}
