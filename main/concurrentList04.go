package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func printPrime(prefix string) {
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}

func main() {
	// 给每个可用的 CPU 分配一个逻辑处理器
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(2)

	go printPrime("A")
	go printPrime("B")

	wg.Wait()

}
