package tcp

import (
	"io"
	"net"
	"net/http"
	"testing"
)

func handleError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("failed: ", err)
	}
}

func TestConn(t *testing.T) {
	lis, err := net.Listen("tcp", ":0")
	handleError(t, err)
	defer lis.Close()

	http.HandleFunc("/hello", helloHandler)
	go http.Serve(lis, nil)

	resp, err := http.Get("http://" + lis.Addr().String() + "/hello")
	handleError(t, err)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	handleError(t, err)

	if string(body) != "hello world!" {
		t.Fatal("expected hello world, but got ", string(body))
	}
}
