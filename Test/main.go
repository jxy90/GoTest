package main

import (
	"fmt"
	"sync"
)

func sum(a int, b int) int {
	return a + b
}

type A struct {
	nums []int
	num  int
	str  string
}

func main() {
	cat, dog, fish := make(chan bool), make(chan bool), make(chan bool)
	finish := make(chan bool)
	wg := sync.WaitGroup{}
	go func() {
		for <-cat {
			fmt.Println("cat")
			dog <- true
		}
	}()
	go func() {
		for <-dog {
			fmt.Println("dog")
			fish <- true
		}
	}()
	go func() {
		for <-fish {
			fmt.Println("fish")
			finish <- true
			wg.Done()
		}
	}()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		cat <- true
		<-finish
	}
	wg.Wait()
}
