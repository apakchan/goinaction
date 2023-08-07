package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func intCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 安全对 counter 执行 +1 操作
		atomic.AddInt64(&counter, 1)
		// 当前 goroutine 从线程退出, 放回到队列
		runtime.Gosched()
	}
}

func main() {
	wg.Add(2)

	go intCounter(1)
	go intCounter(2)

	wg.Wait()

	fmt.Printf("ultimate value: %d\n", counter)
}
