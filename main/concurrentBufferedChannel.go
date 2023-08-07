package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	channelBuffer    = 10 // 工作数量(通道缓冲)
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func worker(tasks chan string, id int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			// 通道空了且被关闭
			fmt.Printf("Worker: %d : Shutting Down\n", id)
			return
		}

		fmt.Printf("Worker %d: Starterd %s\n", id, task)
		time.Sleep(50 * time.Millisecond)

		fmt.Printf("Worker %d: Completed %s\n", id, task)
	}
}

func main() {
	tasks := make(chan string, channelBuffer)

	wg.Add(numberGoroutines)

	for gr := 1; gr <= 4; gr++ {
		go worker(tasks, gr)
	}

	for i := 1; i <= channelBuffer; i++ {
		tasks <- fmt.Sprintf("Task: %d", i)
	}

	// 所有工作处理完关闭通道
	close(tasks)

	wg.Wait()
}
