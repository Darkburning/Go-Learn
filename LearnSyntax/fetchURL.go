package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// 如果url缺少http:/这个前缀为它加上
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err) //%v为任意类型的占位符
			// Fprintf第一个参数需要指定输出的目标，可以是文件、网络连接等任何实现了io.Writer 接口的对象
			// 此处将输出重定向至标准错误流
			os.Exit(1)
		}
		//获取HTTP的报文主体
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil { //如将前一句放在同一行则是局部变量
			// 我们使用了 io.Copy 函数将响应体 resp.Body 中的数据复制到标准输出流 os.Stdout 中，并通过错误处理机制捕获了 io.Copy 函数可能返回的错误
			fmt.Fprintf(os.Stderr, "fetch: reading: %v\n", err)
			os.Exit(1)
		}
		resp.Body.Close() //防止资源泄露

		//获取HTTP报文的响应状态码
		stateCode := resp.Status
		fmt.Println(stateCode)
	}
}
