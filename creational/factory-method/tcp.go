package factorymethod

import "fmt"

type TcpProxy struct {
}

func (pxy *TcpProxy) Run() {
	fmt.Println("TcpProxy:Run")
}

func (pxy *TcpProxy) Close() {
	fmt.Println("TcpProxy:Close")
}

type TcpProxyFactory struct {
}

func (factory *TcpProxyFactory) NewProxy() Proxy {
	fmt.Println("TcpProxyFactory:NewProxy")
	return &TcpProxy{}
}

func NewTcpProxyFactory() ProxyFactory {
	fmt.Println("NewTcpProxyFactory")
	return &TcpProxyFactory{}
}
