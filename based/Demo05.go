package main

import (
	"fmt"
	"sync/atomic"
	"time"
)


func main() {
	var cnt uint32 = 0
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(time.Millisecond)
			atomic.AddUint32(&cnt, 1)
		}()
	}
	time.Sleep(time.Second)
	cntFinal := atomic.LoadUint32(&cnt)
	fmt.Println(cntFinal)
}