package proxy

import (
	"fmt"
	"testing"
)

// go test -v ./structural/proxy/...

func TestNewNginxServer(t *testing.T) {
	server := NewNginxServer()

	code, data := server.ServeHTTP("/")
	fmt.Println(code, data)

	code, data = server.ServeHTTP("/")
	fmt.Println(code, data)

	code, data = server.ServeHTTP("/status")
	fmt.Println(code, data)

	code, data = server.ServeHTTP("/status")
	fmt.Println(code, data)

	code, data = server.ServeHTTP("/not-found")
	fmt.Println(code, data)
}
