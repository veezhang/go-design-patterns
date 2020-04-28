package factorymethod

import "testing"

// go test -v ./creational/factory-method/...

func TestNewProxy(t *testing.T) {
	var proxy Proxy
	var proxyFactory ProxyFactory

	proxyFactory = NewTcpProxyFactory()
	proxy = proxyFactory.NewProxy()
	proxy.Run()
	proxy.Close()
	proxy = proxyFactory.NewProxy()
	proxy.Run()
	proxy.Close()

	proxyFactory = NewUdpProxyFactory()
	proxy = proxyFactory.NewProxy()
	proxy.Run()
	proxy.Close()
	proxy = proxyFactory.NewProxy()
	proxy.Run()
	proxy.Close()
}
