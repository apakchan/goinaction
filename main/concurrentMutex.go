package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	wg      sync.WaitGroup
	mutex   sync.Mutex
	counter int
)

func intCounter(id int) {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		mutex.Lock()
		{
			value := counter

			// 将当前 goroutine 线程退出, 放回等待队列
			runtime.Gosched()

			value++

			counter = value
		}
		mutex.Unlock()
	}
}

func main() {
	wg.Add(2)

	go intCounter(1)
	go intCounter(2)

	wg.Wait()
	fmt.Println("final counter: ", counter)
}
