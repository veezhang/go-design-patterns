package singleton

import (
	"sync"
)

// 在  go 源码 golang/go/src/net/textproto/reader.go 中 commonHeader 也是有使用单例模式的

type singleton struct {
	// 如果是空结构体，编译器会优化，每次 new 都是同一个
	// 所以这里加一个成员
	data int
}

var instance struct {
	once sync.Once
	sg   *singleton
}

func GetInstance() *singleton {
	instance.once.Do(func() {
		instance.sg = new(singleton)
	})
	return instance.sg
}
