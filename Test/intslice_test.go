package main

import (
	"fmt"
	"sort"
	"testing"
)

func Test_IntSlice(t *testing.T) {
	nums := []int{1, 9, 3, 5, 8, 3, 6, 8, 8, 4, 6, 2}
	slice := sort.IntSlice{}
	h := new(hp)
	for _, num := range nums {
		slice = append(slice, num)
		h.Push(num)
	}
	fmt.Println(slice)

	for i := range nums {
		fmt.Println(i)
		fmt.Println(h.Pop())
	}
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
