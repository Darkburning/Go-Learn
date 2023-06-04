package protocol

/*
	Request包含序列号、方法名和参数
	Response包含序列号、错误信息和返回值
*/

type Request struct {
	Method string
	Args   interface{}
}

type Response struct {
	Err     string
	Replies interface{}
}
