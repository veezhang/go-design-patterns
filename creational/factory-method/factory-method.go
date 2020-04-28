package factorymethod

// 摘抄自  frp 源码 : https://github.com/fatedier/frp/blob/v0.33.0/server/proxy
// 这里做了简化，原本是简单工厂模式的，这里改成工厂方法模式。

// 比如我们需要再添加一个 HttpProxy
// 我们需要添加：
//
// 然后添加 http.go ，在其中添加 type HttpProxy struct {} 和 type TcpProxyFactory struct {} 并实现相关方法
// 不再需要修改 porxy.go ，对原有代码没有侵入
//
