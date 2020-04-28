package factorymethod

type ProxyType int

const (
	ProxyTypeTcp ProxyType = iota
	ProxyTypeUdp
)

type Proxy interface {
	Run()
	Close()
}

type ProxyFactory interface {
	NewProxy() Proxy
}
