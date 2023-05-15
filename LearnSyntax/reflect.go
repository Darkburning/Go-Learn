package main

import (
	"log"
	"reflect"
	"strings"
	"sync"
)

func main() {
	var mutex sync.Mutex
	typ := reflect.TypeOf(&mutex)
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		argv := make([]string, 0, method.Type.NumIn())
		returns := make([]string, 0, method.Type.NumOut())
		// 第0个入参是本身
		for j := 1; j < method.Type.NumIn(); j++ {
			argv = append(argv, method.Type.In(j).Name())
		}
		for j := 0; j < method.Type.NumOut(); j++ {
			returns = append(returns, method.Type.Out(j).Name())
		}
		log.Printf("func (m *%s) %s(%s) %s",
			typ.Elem().Name(),
			method.Name,
			strings.Join(argv, ","), // 若有多个参数用逗号拼接
			strings.Join(returns, ","))
	}
}
