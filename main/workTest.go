package main

import (
	"../worker"
	"log"
	"sync"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"jason",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	p := worker.New(8)
	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}
			go func() {
				// 将 np 提交到 Pool 中等待调用
				p.Run(&np)
				wg.Done()
			}()

		}
	}

	wg.Wait()
	p.Shutdown()
}
