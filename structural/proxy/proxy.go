package proxy

// Proxy 可能大家很容易就想到了 nginx 代理
// 这里模拟 nginx proxy 来了解 Proxy 代理模式

// 为了子一个文件中直观的看，这里不放到不同的文件

////////////////////////////////////////////////////////////////////////////////
// handler.go

// Handler 类比： 抽象主题（Subject）类
// RealServer 和 NginxServer 都会实现此接口
type Handler interface {
	ServeHTTP(url string) (int, string)
}

////////////////////////////////////////////////////////////////////////////////
// nginx-server.go
func NewNginxServer() *NginxServer {
	return &NginxServer{
		Server:            &RealServer{},
		MaxAllowedRequest: 1,
		RateLimiter:       make(map[string]int),
	}
}

// NginxServer 类比： 代理（Proxy）类
type NginxServer struct {
	Server            *RealServer
	MaxAllowedRequest int
	RateLimiter       map[string]int
}

func (s *NginxServer) ServeHTTP(url string) (int, string) {
	if !s.checkRateLimiter(url) {
		return 403, "Not Allowed"
	}
	// 这里还可以做其他预处理
	code, data := s.Server.ServeHTTP(url)
	// 这里还可以做其他后处理
	return code, data
}

// 简单的一个限制
func (s *NginxServer) checkRateLimiter(url string) bool {
	n := s.RateLimiter[url] // 如果不存在， 默认值是 0
	n++
	s.RateLimiter[url] = n

	if n > s.MaxAllowedRequest {
		return false
	}
	return true
}

////////////////////////////////////////////////////////////////////////////////
// real-server.go

// RealServer 类比： 真实主题（Real Subject）类
type RealServer struct {
}

func (s *RealServer) ServeHTTP(url string) (int, string) {
	if url == "/" {
		return 200, "Hello Proxy"
	}

	if url == "/status" {
		return 200, "OK"
	}

	return 404, "Not Found"
}
