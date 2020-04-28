package simplefactory

import "fmt"

type TcpProxy struct {
}

func (pxy *TcpProxy) Run() {
	fmt.Println("TcpProxy:Run")
}

func (pxy *TcpProxy) Close() {
	fmt.Println("TcpProxy:Close")
}
