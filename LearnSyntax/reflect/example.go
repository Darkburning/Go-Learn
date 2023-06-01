package main

import (
	"log"
	"reflect"
	"strings"
)

func main() {
	log.SetFlags(0)
	//反射进阶用法：
	var reflectValue reflect.Value
	// sync.Mutex的接受者为指针接受者
	typ := reflect.TypeOf(&reflectValue)
	// 迭代这个struct的所有方法，拿到参数与返回值
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
