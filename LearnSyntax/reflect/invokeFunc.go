package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
}

func (s *Student) DoHomework(num int) {
	fmt.Printf("%s is doing homework %d", s.name, num)
}
func main() {
	s := Student{name: "Kali"}
	v := reflect.ValueOf(&s)
	vMethod := v.MethodByName("DoHomework")
	if vMethod.IsValid() {
		// 入参必须是reflect.Value切片
		in := []reflect.Value{reflect.ValueOf(55)}
		vMethod.Call(in)
	}
}
