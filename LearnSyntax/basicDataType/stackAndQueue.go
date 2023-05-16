package main

import "fmt"

func main() {
	stack := make([]int, 0)
	stack = append(stack, 2)
	stack = append(stack, 3)
	stack = append(stack, 4)
	fmt.Println(stack)
	stack = stack[:len(stack)-1]
	fmt.Println(stack)
	top := stack[len(stack)-1]
	fmt.Println(top)

	queue := make([]int, 0)
	queue = append(queue, 4)
	queue = append(queue, 5)
	queue = append(queue, 6)
	fmt.Println(queue)
	queue = queue[1:]
	fmt.Println(queue)
	front := queue[0]
	fmt.Println(front)
}
