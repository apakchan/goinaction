package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

// court 就是通道数据
func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court

		if !ok {
			// 通道被关闭, 赢了
			fmt.Printf("Player %s Won!\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("%s missed\n", name)

			// 关闭通道, 表示 name 输了
			close(court)
			return
		}

		fmt.Printf("%s Hit %d\n", name, ball)
		ball++
		// 通道中保存目前是第几发
		court <- ball
	}
}

func main() {
	// 通道数据
	// 无缓冲
	court := make(chan int)
	wg.Add(2)

	go player("A", court)
	go player("B", court)

	// 赋初值 1
	// 一定是这行代码运行完 goroutine 才得以运行
	court <- 1

	wg.Wait()
}
