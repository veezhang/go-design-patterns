package simplefactory

// 摘抄自  frp 源码 : https://github.com/fatedier/frp/blob/v0.33.0/server/proxy
// 这里做了简化，主要体现 简单工厂模式

// 比如我们需要再添加一个 HttpProxy
// 我们需要添加：
//
// 在 porxy.go 中添加：const ProxyTypeHttp， 并且在 NewProxy 中需要添加 case 分支， 这里会入侵代码
// 然后添加 http.go ，在其中添加 type HttpProxy struct {} 并实现相关方法
//
