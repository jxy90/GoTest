package test

import (
	"log"
	"sync"
	"testing"
)

func Test_0(t *testing.T) {
	//fmt.Println(kIncreasing([]int{4, 1, 5, 2, 6, 2}, 3))
}
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func TestFan(t *testing.T) {
	c := gen(2, 3, 4, 5, 6, 7, 8)
	out1 := sq(c)
	out2 := sq(c)

	for n := range merge(out1, out2) {
		log.Print(n)
	}
}
