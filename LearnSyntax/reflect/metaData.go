package main

import (
	"fmt"
	"reflect"
)

func printMeta(obj interface{}) {
	t := reflect.TypeOf(obj)
	n := t.Name()
	k := t.Kind()
	v := reflect.ValueOf(obj)
	fmt.Printf("Type: %s, Name: %s, Kind: %s, Value: %v\n", t, n, k, v)
}

func main() {

	// 反射基础用法
	var a uint64 = 10
	sli := make([]string, 10)
	sum := func(a, b int) int {
		return a + b
	}
	type Cat struct {
		Name string
		Age  uint64
	}
	var Mimi Cat = Cat{
		Name: "Mimi",
		Age:  9,
	}
	type function func(a, b int) int
	var sub function = func(a, b int) int {
		return a - b
	}

	printMeta(a)
	printMeta(sli)
	printMeta(sum)
	printMeta(Mimi)
	printMeta(sub)

}
