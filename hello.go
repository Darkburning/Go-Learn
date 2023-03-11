package main

import (
	"errors"
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
// 在 Go 中，interface 类型表示了一组方法签名，任何实现了这些方法的类型都可以被赋值给该接口变量
type geometry interface { // 定义一个几何体的接口
	area() float32
	peri() float32
}
type rect struct {
	length float32
	width  float32
}

// 相当于实现了rect对象的对应方法
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

// 错误处理
// 错误通常是最后一个返回值并且是 error 类型，它是一个内建的接口
func f1(arg int) (int, error) {
	if arg == 42 {

		return -1, errors.New("can't work with 42")

	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

// 定义了结构体argError的Error方法
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// 注意error是interface
func f2(arg int) (int, error) {
	if arg == 42 {

		return -1, &argError{arg, "can't work with it"} //因为Error需要指针接收
	}
	return arg + 3, nil
}

// 用于go 并发
func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func send(msg chan string, str string) {
	msg <- str
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
	//fmt.Scanf("%d", &i)
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
	//rect1 := rect{2, 3}
	//circle1 := circle{3}
	//fmt.Printf("Area is %f, Peri is %f\n", rect1.area(), rect1.peri())
	//measure(rect1)
	//measure(circle1)
	//
	//// 嵌入（Embedding）
	//co := container{base: base{num: 1},
	//	str: "some name",
	//}
	//// 可直接访问嵌入其内的数据 等效于 fmt.Println("also num:", co.base.num)
	//fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str) //当创建含有嵌入的结构体，必须对嵌入进行显式的初始化
	//// 也可以调用嵌入其中的数据
	//fmt.Println("describe:", co.describe())
	//// 定义一个接口
	//type describer interface {
	//	describe() string
	//}
	//// co的嵌入已经实现接口故可直接调用接口方法
	//var d describer = co
	//fmt.Println("describer:", d.describe())

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

	//for _, i := range []int{7, 42} {
	//	if r, e := f1(i); e != nil { //此处相当于先执行 r, e := f1(i)然后进行判断,这样写更紧凑
	//		fmt.Println("f1 failed:", e)
	//	} else {
	//		fmt.Println("f1 worked:", r)
	//	}
	//}
	//for _, i := range []int{7, 42} {
	//	if r, e := f2(i); e != nil {
	//		fmt.Println("f2 failed:", e)
	//	} else {
	//		fmt.Println("f2 worked:", r)
	//	}
	//}
	//// 由于argError实现了error这个interface的方法因此可以将argError赋值给interface，此时e是一个*argError类型变量
	//_, e := f2(42)
	//if ae, ok := e.(*argError); ok { //此处进行了类型断言，e.(*argError)在判断e是否是*argError类型
	//	fmt.Println(ae.arg) //如果是则将argError赋值给ae并将ok置为true，否则ae将被置为nil，ok置为false
	//	fmt.Println(ae.prob)
	//}

	//go channel
	//默认发送和接收操作是阻塞的，直到发送方和接收方都就绪。这个特性允许我们，不使用任何其它的同步操作,就可以在程序结尾处等待消息
	//var inputMsg string
	//fmt.Scanf("%s", &inputMsg)
	//message := make(chan string) // 使用make(chan val-type)创建一个接受string的channel
	//go send(message, inputMsg)   //将"ping"发送到message这个channel
	//msg := <-message             // 定义string类型的msg接受message的值
	//fmt.Println(msg)

	//通道缓冲
	//默认情况下，通道是无缓冲的，这意味着只有对应的接收（<- chan）
	//通道准备好接收时，才允许进行发送（chan <-）。有缓冲通道 允许在没有对应接收者的情况下，缓存一定数量的值
	message := make(chan string, 2)
	message <- "buffered"
	message <- "channel"
	fmt.Println(<-message) //由于通道是有缓冲的， 因此我们可以将这些值发送到通道中，而无需并发的接收。
	fmt.Println(<-message)
}
