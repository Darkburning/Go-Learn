package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 回撤请求，释放资源
func handler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	// 多路复用
	called := make(chan struct{})
	go func() {
		time.Sleep(4 * time.Second)
		called <- struct{}{}
	}()

	select {
	case <-called:
		fmt.Println("Server Called Success!")
		fmt.Fprintf(w, "Server Called Success!")
	case <-ctx.Done():
		fmt.Println("Client Cancel!")
		fmt.Fprintf(w, "Client Cancel!")
	}
}
func main() {
	// register the handler function
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
