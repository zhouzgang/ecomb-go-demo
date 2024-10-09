package main

import (
	"sync"
	"sync/atomic"
)

/*
what is different between type and var?
type 类似 Java 匿名对象
struct 的使用，为什么这里可以继承 lock 函数
 */
var total struct{
	sync.Mutex
	value int
}

func worker01(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		total.Lock()
		total.value += 1
		total.Unlock()
	}
}

var total02 uint64

func worker02(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		atomic.AddUint64(&total02, uint64(1))
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker01(&wg)
	go worker01(&wg)
	wg.Wait()

	wg.Add(2)
	go worker02(&wg)
	go worker02(&wg)
	wg.Wait()

	println("worker 01: ", total.value)
	println("worker 02: ", total02)
}
