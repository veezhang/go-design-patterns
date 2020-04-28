package simplefactory

import "testing"

// go test -v ./creational/simple-factory/...

func TestNewProxy(t *testing.T) {

	var proxy Proxy

	proxy, _ = NewProxy(ProxyTypeTcp)
	proxy.Run()
	proxy.Close()
	proxy, _ = NewProxy(ProxyTypeTcp)
	proxy.Run()
	proxy.Close()

	proxy, _ = NewProxy(ProxyTypeUdp)
	proxy.Run()
	proxy.Close()
	proxy, _ = NewProxy(ProxyTypeUdp)
	proxy.Run()
	proxy.Close()
}
