//package main
//
//import (
//	"log"
//	"reflect"
//	"strings"
//)
//
//func main() {
//	log.SetFlags(0)
//	//反射进阶用法：
//	var reflectValue reflect.Value
//	// sync.Mutex的接受者为指针接受者
//	typ := reflect.TypeOf(&reflectValue)
//	// 迭代这个struct的所有方法，拿到参数与返回值
//	for i := 0; i < typ.NumMethod(); i++ {
//		method := typ.Method(i)
//
//		argv := make([]string, 0, method.Type.NumIn())
//		returns := make([]string, 0, method.Type.NumOut())
//		// 第0个入参是本身
//		for j := 1; j < method.Type.NumIn(); j++ {
//			argv = append(argv, method.Type.In(j).Name())
//		}
//		for j := 0; j < method.Type.NumOut(); j++ {
//			returns = append(returns, method.Type.Out(j).Name())
//		}
//		log.Printf("func (m *%s) %s(%s) %s",
//			typ.Elem().Name(),
//			method.Name,
//			strings.Join(argv, ","), // 若有多个参数用逗号拼接
//			strings.Join(returns, ","))
//	}
//}

package main

import (
	"fmt"
	"reflect"
)

func foo(x int, y string) {
	fmt.Printf("foo(%d, %s)\n", x, y)
}

func main() {
	// 获取函数类型及其入参类型
	ft := reflect.TypeOf(foo)
	fmt.Printf("Function type: %v\n", ft)

	// 遍历入参类型并输出相关信息
	for i := 0; i < ft.NumIn(); i++ {
		in := ft.In(i)
		fmt.Printf("Input %d: name=%s, type=%v, kind=%v\n", i, in.Name(), in.String(), in.Kind())
	}
}
