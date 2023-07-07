package main

import (
	"fmt"
	"time"
	"unsafe"
)

type EST struct {
}

func main() {
	var a struct{}
	var b EST

	fmt.Printf("addr: %p, size: %d\n", &a, unsafe.Sizeof(a))
	fmt.Printf("addr: %p, size: %d\n", &b, unsafe.Sizeof(b))

	// 实现Set
	Set := make(map[string]struct{}, 10)
	Set["Hello"] = struct{}{}
	Set["World"] = EST{}
	Set["!"] = struct{}{}
	fmt.Printf("len: %d\n", len(Set))

	ch := make(chan struct{})
	go func(chan struct{}) {
		time.Sleep(time.Second * 2)
		fmt.Println("goroutine exit")
		ch <- struct{}{}
	}(ch)
	<-ch
}
