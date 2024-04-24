package _024

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_baidu(t *testing.T) {
	pool := NewWorkerPool(5)

	for i := 0; i < 10; i++ {
		pool.Run(func() {
			fmt.Println("Task:", i)
			time.Sleep(time.Second * 1)
		})
	}

	pool.Shutdown()
}

//go 协程池
type WorkerPool struct {
	work chan func()
	wg   sync.WaitGroup
}

func NewWorkerPool(maxWorkers int) *WorkerPool {
	pool := &WorkerPool{work: make(chan func(), maxWorkers)}
	for i := 0; i < maxWorkers; i++ {
		pool.wg.Add(1)
		go pool.worker()
	}
	return pool
}

func (pool *WorkerPool) worker() {
	defer pool.wg.Done()
	for job := range pool.work {
		job()
	}
}

func (pool *WorkerPool) Run(task func()) {
	pool.work <- task
}

func (pool *WorkerPool) Shutdown() {
	close(pool.work)
	pool.wg.Wait()
}
