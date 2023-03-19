package main

import "fmt"

type degree float32

// 定义一个String方法，因为当使用fmt包的打印方法时，将会优先使用该类型对应的String方法返回的结果打印
func (d degree) String() string {
	return fmt.Sprintf("%.2f°", d)
}
func k(list []string) string { return fmt.Sprintf("%q", list) } //如实打印
func main() {
	const d degree = 15.000003
	fmt.Printf("%v\n", d)
	var s []string
	s = append(s, "Hello")
	s = append(s, "World!")
	fmt.Println(s)
	str := k(s)
	fmt.Println(str)
	// 打印函数签名或值的类型
	fmt.Printf("%T\t%T", k, d)
}
