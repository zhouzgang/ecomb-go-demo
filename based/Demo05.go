package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var totalTickets int32 = 10
var mutex = &sync.Mutex{}

func sellTickets(i int) {
	for totalTickets > 0 {
		mutex.Lock()
		if totalTickets > 0 {
			time.Sleep(time.Duration(rand.Intn(50)) * time.Microsecond)
			totalTickets--
			fmt.Println("id:", i, "tickets: ", totalTickets)
		}
		mutex.Unlock()
	}
}

func main() {
	runtime.GOMAXPROCS(4)
	rand.Seed(time.Now().Unix())

	for i := 0; i < 5; i++ {
		go sellTickets(i)
	}

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}