package main

import (
	"fmt"
	"sync"
)

var cnt = 0

func Increment(group *sync.WaitGroup) {
	defer group.Done()
	for i := 0; i < 100000; i++ {
		cnt++
	}
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go Increment(wg)
	go Increment(wg)
	wg.Wait()
	fmt.Printf("cnt: %d", cnt)

}
