package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool 管理一组可以安全在多个 goroutine 间共享的资源。被管理的资源必须实现 io.Closer 接口
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// ErrPoolClosed 请求一个已经关闭的池
var ErrPoolClosed = errors.New("Pool has been closed.")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire 从池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// 检查是否有空闲资源
	case r, ok := <-p.resources:
		log.Println("Acquire: ", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	// 没有空闲资源, 提供新资源
	default:
		log.Println("Acquire:", "New Resources")
		return p.factory()
	}
}

// Release 将一个使用后的资源放回池中
func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	// 将这个资源放入队列
	case p.resources <- r:
		log.Println("Release:", "In Queue")
	// 如果队列已满, 关闭这个资源
	default:
		log.Println("Release", "Closing")
		r.Close()
	}
}

func (p *Pool) Close() {
	// 保证 Close 与 Release 互斥
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	// 关闭 resources 通道
	close(p.resources)

	// 清除 resources 通道中的资源
	for resource := range p.resources {
		resource.Close()
	}
}
