package main

import (
	"fmt"
	"time"
)

type f func(arr [10]int, size int) //声明一个类型

func display(arr [10]int, size int) {
	for i := 0; i < size; i++ {
		fmt.Println(arr[i])
	}
}
func initArr(arr [10]int, size int, show f) { //函数作为参数被传入函数
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	show(arr, size)
}
func show(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func say(s string) { // 用于go 并发
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func fib(n int) int { // 斐波那契数列
	var arr []int // 利用slice实现，数组不接受非常数的size
	arr = append(arr, 0, 1)
	for i := 2; i <= n; i++ {
		arr = append(arr, arr[i-1]+arr[i-2])
	}
	return arr[n]
}

func main() { // go对大括号的放置十分严格，仅允许这样放置
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World!")

	//// 在go中声明时指定大小的为数组，否则为切片slice
	//var a [10]int //go数组
	////initArr(a, 8, display)
	//for i := 0; i < 10; i++ {
	//	a[i] = i
	//}

	//s := []int{1, 2, 3} // slice动态数组
	//s = append(s, 4)
	//s = append(s, 4, 5, 6)
	////s := a[:] //slice支持切片截取
	//test := s
	//test = append(test, 100)
	//show(test)
	//fmt.Print(reflect.TypeOf(test))
	//fmt.Print(len(s), cap(s), "\n") // len() 和 cap()显示slice的长度和容量

	//var b int
	//fmt.Println("Input a number to compute fib:")
	//fmt.Scanf("%d", &b)
	//fmt.Println(fib(b))
	//fmt.Printf("sin(1) is: %f \n", math.Sin(1))

	type Person struct { //声明一个结构体
		name string
		sex  string
		age  int
	}
	// 结构体对象构造方式
	P1 := Person{"JAY", "Male", 18}
	fmt.Println(P1)

	// map 构造方法
	myMap := make(map[string]string)
	myMap["red"] = "r"
	myMap["yellow"] = "y"
	myMap["blue"] = "b"
	for i, j := range myMap {
		fmt.Printf("%s:%s\n", i, j)
	}
	delete(myMap, "blue") //删除键值对
	for i, j := range myMap {
		fmt.Printf("%s:%s\n", i, j)
	}

	//go say("Hello") //go 并发
	//go say("World")
	//say("!")
}
