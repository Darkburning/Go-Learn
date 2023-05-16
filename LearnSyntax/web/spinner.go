package main

import (
	"fmt"
	"time"
)

func Fib(n int) int {
	if n < 2 {
		return n
	} else {
		return Fib(n-1) + Fib(n-2)
	}
}

// 并发执行可产生动画
func spinner(delay time.Duration) {
	for {
		for _, v := range `-\|/-` { //此处使用反引号可避免转义且表明这是一个字符串常量
			fmt.Printf("\r%c", v)
			time.Sleep(delay)
		}
	}
}
func main() {
	const times = 45
	go spinner(100 * time.Millisecond)
	fmt.Printf("\rFib(%d) = %d", times, Fib(times))
}
