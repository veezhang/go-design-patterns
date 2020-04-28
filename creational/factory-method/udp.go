package factorymethod

import "fmt"

type UdpProxy struct {
}

func (pxy *UdpProxy) Run() {
	fmt.Println("UdpProxy:Run")
}

func (pxy *UdpProxy) Close() {
	fmt.Println("UdpProxy:Close")
}

type UdpProxyFactory struct {
}

func (factory *UdpProxyFactory) NewProxy() Proxy {
	fmt.Println("UdpProxyFactory:NewProxy")
	return &UdpProxy{}
}

func NewUdpProxyFactory() ProxyFactory {
	fmt.Println("NewUdpProxyFactory")
	return &UdpProxyFactory{}
}
