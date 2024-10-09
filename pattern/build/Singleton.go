package main

import (
	"sync"
	"sync/atomic"
)


// 一个空的匿名对象？
// 这种空的匿名对象是什么时候创建的，
type singleton struct {}

var (
	instance *singleton
	initialzed uint32
	mu sync.Mutex
)

// 这里会返回一个包私有类型的指针，但是还不知道这样有什么不好的
func Instance() *singleton {
	if atomic.LoadUint32(&initialzed) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialzed, 1)
		instance = &singleton{}
	}
	return instance
}

func main() {

}
