package main

import (
	"context"
	"fmt"
	"time"
)

func startServer(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Client Cancel!")
	case <-time.After(time.Second * 2):
		fmt.Println("Server Done Success")
	}

}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	startServer(ctx)

}
