package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
}

func (s *Student) DoHomework(num int) {
	fmt.Printf("%s is doing homework %d\n", s.name, num)
}

func sum(a, b int) int {
	return a + b
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

	fun := reflect.ValueOf(sum)
	inArgs := []reflect.Value{reflect.ValueOf(2), reflect.ValueOf(3)}
	fmt.Printf("%s\n", inArgs)
	ret := fun.Call(inArgs)
	fmt.Printf("sum: %d", ret[0].Interface())

}
