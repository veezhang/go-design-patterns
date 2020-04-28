package simplefactory

import "fmt"

type ProxyType int

const (
	ProxyTypeTcp ProxyType = iota
	ProxyTypeUdp
)

type Proxy interface {
	Run()
	Close()
}

func NewProxy(proxyType ProxyType) (pxy Proxy, err error) {
	switch proxyType {
	case ProxyTypeTcp:
		pxy = &TcpProxy{}
	case ProxyTypeUdp:
		pxy = &UdpProxy{}
	default:
		return pxy, fmt.Errorf("proxy type not support")
	}
	return
}
