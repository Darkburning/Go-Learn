package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("请输入你想写入的内容:")
	var content string
	//创建一个bufio.Reader对象，并将其与标准输入流(os.Stdin)关联。接着使用ReadString('\n')方法从标准输入流中读取一行输入，
	//直到遇到换行符为止，读取到的内容包括换行符
	reader := bufio.NewReader(os.Stdin)
	content, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	// writeFile(content)
	appendFile(content)
	readFile()
}

func appendFile(content string) {
	for _, file := range os.Args[1:] {
		// 文件设置为APPEND模式、没有文件则创建、只能写
		f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		//defer f.Close() // 将该函数的执行推迟到所在函数执行完毕之后，此处避免因为错误或异常导致资源泄露
		n, err := f.WriteString(content)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("TotalNum is:%v\n", n)
		f.Close()
	}

}

func writeFile(content string) {
	// 依次读取命令行参数打开文件往里写入string对应的内容
	for _, file := range os.Args[1:] {
		os.WriteFile(file, []byte(content), 0777) // 如果没有文件会创建文件，如果有会覆盖
	}
}
func readFile() {
	for _, file := range os.Args[1:] {
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err) //%v为任意类型的占位符
			// Fprintf第一个参数需要指定输出的目标，可以是文件、网络连接等任何实现了io.Writer 接口的对象
			// 此处将输出重定向至标准错误流
			os.Exit(1)
		}
		//在Printf函数中，%q参数会将字节数组中的每个字节转换为ASCII码对应的字符，如果字节无法转换为可打印字符，则会使用转义序列进行表示
		fmt.Printf("%q\n", content)
	}

}
