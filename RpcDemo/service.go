package main

// 定义服务

func Increment(args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func Multiply(args ...int) int {
	product := 1
	for i := 0; i < len(args); i++ {
		product *= args[i]
	}
	return product
}
