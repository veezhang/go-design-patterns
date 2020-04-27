package prototype

// 摘抄自 go 源码 ： golang/go/src/net/http/header.go
// 这里做了简化，主要体现 原型模式

// 其它的kubernetes 也使用了原型模式
// github/kubernetes/kubernetes/pkg/kubelet/cm/cpuset/cpuset.go
// func (s CPUSet) Clone() CPUSet { ... }

type Header map[string]string

// Set sets the value of key
func (h Header) Set(key, value string) {
	h[key] = value
}

// Set returns the value of key
func (h Header) Get(key string) string {
	if v, ok := h[key]; ok {
		return v
	}
	return ""
}

func (h Header) Clone() Header {
	if h == nil {
		return nil
	}

	h2 := make(Header, len(h))
	for k, v := range h {
		h2[k] = v
	}

	return h2
}
