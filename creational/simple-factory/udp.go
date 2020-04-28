package simplefactory

import "fmt"

type UdpProxy struct {
}

func (pxy *UdpProxy) Run() {
	fmt.Println("UdpProxy:Run")
}

func (pxy *UdpProxy) Close() {
	fmt.Println("UdpProxy:Close")
}
