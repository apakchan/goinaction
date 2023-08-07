package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	wg sync.WaitGroup

	shutdown int64
)

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s\n", name)
			break
		}
	}
}

func main() {
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	// run doWork
	time.Sleep(1 * time.Second)

	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()
}
