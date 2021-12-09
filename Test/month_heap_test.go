package main

import (
	"container/heap"
	"fmt"
	"testing"
)

func Test_MonthHeap(t *testing.T) {
	h := &MonthHeap{"Feb", "Mar"}
	heap.Init(h)
	heap.Push(h, "May")
	heap.Push(h, "Apr")
	heap.Push(h, "Jan")
	// first month: Jan
	fmt.Println("first month:", (*h)[0])

	// 输出: Jan	Feb	Mar	Apr	May
	for h.Len() > 0 {
		fmt.Printf("%s\t", heap.Pop(h)) // 注意不是h.Pop()
	}
	fmt.Println()
}
