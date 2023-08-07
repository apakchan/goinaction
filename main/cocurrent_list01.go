package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 分配程序逻辑处理器数量
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	// 表示要等待 2 个 goroutine
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		// 函数退出时调用 Done 通知 main 函数工作完成
		// defer 表示函数返回时调用的代码
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	fmt.Println("Waiting To Finish")
	// 等待 goroutine 结束
	// 如果 Waitgroup 值大于 0, Wait 会阻塞
	wg.Wait()

	fmt.Println("Terminating Program")
}
