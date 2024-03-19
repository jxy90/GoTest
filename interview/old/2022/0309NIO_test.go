package _022

import (
	"fmt"
	"sync"
	"testing"
)

func Test_0309NIO(t *testing.T) {
	//fmt.Println("123")
	ma()
}

//1~10,10个goroutine
var i = 0
var mu sync.Mutex
var wg sync.WaitGroup

func ma() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func inc() {
	mu.Lock()
	defer mu.Unlock()
	i++
	fmt.Println(i)
	wg.Done()
}
