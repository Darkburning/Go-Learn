package main

import (
	"fmt"
	"math"
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
func show(arr []int) { //其实用fmt.Println()即可打印数组或slice
	for i := 0; i < len(arr); i++ {
		fmt.Printf(" %d", arr[i])
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

// go语言支持函数有多个返回值（类似python）
func multipleRtv(a float32, b float32) (plus, minus, multiply, divide float32) {
	return a + b, a - b, a * b, a - b
}

// go支持变参函数
func sum(nums ...int) {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	fmt.Println(totalSum)
}

// 匿名函数和闭包
func intSeq() func() int { //返回值为：一个返回整型变量的函数
	i := 0
	return func() int {
		i++
		return i
	}
}

// 结构体、方法及接口
type geometry interface { // 定义一个几何体的接口
	area() float32
	peri() float32
}
type rect struct {
	length float32
	width  float32
}

// 这里的 area 是一个拥有 *rect 类型接收器(receiver)的方法
func (r rect) area() float32 { // 这里若带*号表示引用传递，无需再拷贝一次，且在函数内可改变对象的值，不带*号表示按值传递
	return r.length * r.width
}
func (r rect) peri() float32 {
	return 2 * (r.length + r.width)
}

// 要在 Go 中实现一个接口，我们只需要实现接口中的所有方法
//	这里我们为 rect 和 circle实现了 geometry 接口

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}
func (c circle) peri() float32 {
	return 2 * math.Pi * c.radius
}

// 如果一个变量实现了某个接口，我们就可以调用指定接口中的方法。
// 这儿有一个通用的 measure 函数，我们可以通过它来使用所有的geometry
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Peri is ", g.peri())
	fmt.Println("Area is ", g.area())
}

type base struct {
	num int
}

// 定义了base的方法
func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// container嵌入（Embedding）base
type container struct {
	base
	str string
}

// 用于go 并发
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() { // go对大括号的放置十分严格，仅允许这样放置
	fmt.Println("Hello, World!")
	// go字符串
	//const s1 = "Hello, World!" //注意const变量的声明方式
	//const s2 = "你好世界!"
	//s3 := "你好Go！"
	//fmt.Println(reflect.TypeOf(s1), reflect.TypeOf(s2), reflect.TypeOf(s3))
	//fmt.Println(len(s1), len(s2), len(s3))
	//for i := 0; i < len(s2); i++ {
	//	fmt.Printf("%x  ", s2[i])
	//}
	//fmt.Println("Rune count:", utf8.RuneCountInString(s2))
	//for idx, runeValue := range s2 {
	//	fmt.Printf("%U starts at %d\n", runeValue, idx)
	//}

	//go if else的变态要求,else必须和if的右括号在同一行，if条件必须与左括号同一行
	//num := 0
	//if num == 0 {
	//	fmt.Println(num + 1)
	//} else {
	//	fmt.Println(num - 1)
	//}

	// switch and for loop(while lop)
	//for true {
	//	var i int
	//	fmt.Println("Input a num:")
	//	fmt.Scanf("%d", &i)
	//	switch i {
	//	case 1:
	//		fmt.Printf("%d is one\n", i)
	//	case 2:
	//		fmt.Printf("%d is two\n", i)
	//	case 3:
	//		fmt.Printf("%d is three\n", i)
	//	case 4:
	//		fmt.Printf("%d is four\n", i)
	//	case 5:
	//		fmt.Printf("%d is five\n", i)
	//	default: // 可选
	//		fmt.Printf("%d is not between 1 and 5\n", i)
	//	}
	//}

	//// 在go中声明时指定大小的为数组，否则为切片slice
	//var a [10]int //go数组
	////initArr(a, 8, display)
	//for i := 0; i < 10; i++ {
	//	a[i] = i
	//}

	// slice动态数组声明与初始化
	//s := make([]int, 3) //等效于s := []int{1, 2, 3}
	//s[0] = 1
	//s[1] = 2
	//s[2] = 3
	// slice相关方法
	//s = append(s, 4) // append会返回一个新的slice
	//s = append(s, 4, 5, 6)
	////s := a[:] //slice支持切片截取
	//test := make([]int, len(s))
	//copy(test, s) //将s复制到test
	//show(test)
	//test = append(test, 100)
	//show(test)
	//fmt.Print(reflect.TypeOf(test))
	//fmt.Print(len(s), cap(s), "\n") // len() 和 cap()显示slice的长度和容量
	//slice作为多维数据结构，每维大小可不同
	//twoD := make([][]int, 3) //注意此处声明二维数组的时候用 :=
	//for i := 0; i < len(twoD); i++ {
	//	innerLen := i + 1
	//	twoD[i] = make([]int, innerLen) //此处指定里面的长度只用=
	//	for j := 0; j < innerLen; j++ {
	//		twoD[i][j] = i + j
	//	}
	//}
	//fmt.Println("2D: ", twoD)

	//var b int
	//fmt.Println("Input a number to compute fib:")
	//fmt.Scanf("%d", &b)
	//fmt.Println(fib(b))
	//fmt.Printf("sin(1) is: %f \n", math.Sin(1))

	// 结构体、方法及接口
	rect1 := rect{2, 3}
	circle1 := circle{3}
	fmt.Printf("Area is %f, Peri is %f\n", rect1.area(), rect1.peri())
	measure(rect1)
	measure(circle1)

	// 嵌入（Embedding）
	co := container{base: base{num: 1},
		str: "some name",
	}
	// 可直接访问嵌入其内的数据 等效于 fmt.Println("also num:", co.base.num)
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str) //当创建含有嵌入的结构体，必须对嵌入进行显式的初始化
	// 也可以调用嵌入其中的数据
	fmt.Println("describe:", co.describe())
	// 定义一个接口
	type describer interface {
		describe() string
	}
	// co的嵌入已经实现接口故可直接调用接口方法
	var d describer = co
	fmt.Println("describer:", d.describe())

	//// map 构造方法
	//myMap := make(map[string]int)
	//myMap["red"] = 1
	//myMap["yellow"] = 2
	//myMap["blue"] = 3
	//for i, j := range myMap { //range遍历
	//	fmt.Printf("%s:%d\n", i, j)
	//}
	//// 当从一个 map 中取值时，还有可以选择是否接收的第二个返回值
	////该值表明了 map 中是否存在这个键。
	////这可以用来消除键不存在 和 键的值为零值 产生的歧义
	//_, isExisted := myMap["red"] // isExisted 第一次声明并赋值需要用:=
	//fmt.Println(isExisted)
	//_, isExisted = myMap["brown"] // isExisted 第二次赋值用=即可
	//fmt.Println(isExisted)
	//
	//delete(myMap, "blue") //删除键值对
	//for i, j := range myMap {
	//	fmt.Printf("%s:%d\n", i, j)
	//}

	//multipleRtv函数多返回值
	//fmt.Println("Input two nums:")
	//var a, b float32
	//fmt.Scanf("%f%f", &a, &b)
	//// 如果不需要某些返回值可用_作为占位符
	//plus, minus, multiply, divide := multipleRtv(a, b)
	//// Printf函数和c语言的差不多
	//fmt.Printf("%.2f + %.2f= %.2f\n", a, b, plus)
	//fmt.Printf("%f - %f= %f\n", a, b, minus)
	//fmt.Printf("%f * %f= %f\n", a, b, multiply)
	//fmt.Printf("%f / %f= %f\n", a, b, divide)

	// 变参函数
	//sum(1, 2, 3)
	//sum(1, 2, 3, 4)
	//nums := []int{1, 2, 3, 4, 5}
	////如果要将数组或slice作为参数需要采用...作为解包提示符
	//sum(nums...)
	// 匿名函数&闭包
	// 匿名函数用于定义一个不需要命名的内联函数
	// 随时随地可写，主要实现一些逻辑简单的函数
	//f := intSeq() // intSeq将一个匿名函数返回给f

	//fmt.Println(f()) //由于f是一个函数，调用时要+()
	//fmt.Println(f())
	//fmt.Println(f())
	//t := intSeq()
	//fmt.Println(t())

	//go 并发
	//go say("Hello")
	//go say("World")
	//say("!")
}
