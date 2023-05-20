package main

import (
	"context"
	"log"
	"time"
)

func doSomething(ctx context.Context) {
	// 通过select chan的方式来实现回撤请求，释放资源
	select {
	case <-time.After(2 * time.Second):
		log.Println("finish doingSomething")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(err.Error())
	}
}

func main() {
	// 利用withCancel方法得到一个cancelFunction
	ctx, cancel := context.WithCancel(context.Background())

	// 可理解为用户等待一秒后取消
	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	doSomething(ctx)
}
