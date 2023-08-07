package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	// interrupt 通道报告从 os 发送的信号
	interrupt chan os.Signal

	// complete 通道报告处理任务完成
	complete chan error

	// timeout 报告处理任务超时
	timeout <-chan time.Time

	// 持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

var ErrTimeout = errors.New("received timeout")

var ErrInterrupt = errors.New("received interrupt")

// New Factory Method
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add append tasks to r.tasks
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start run r.tasks
func (r *Runner) Start() error {
	// 接受所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	// tasks 运行时异常
	case err := <-r.complete:
		return err
	// 超时异常
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt // 返回中断异常
		}
		// run task
		task(id)
	}

	return nil
}

// 验证是否收到中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	// 中断事件被触发发出的信号
	case <-r.interrupt:
		// 通知接受后续的任何信号
		signal.Stop(r.interrupt)
		return true
	// 继续正常运行
	default:
		return false
	}
}
