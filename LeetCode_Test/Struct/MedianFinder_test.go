package Struct

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

func Test_LRUCache(t *testing.T) {
	obj := ConstructorMedianFinder()
	obj.AddNum(-1)
	obj.AddNum(-2)
	param_2 := obj.FindMedian()
	fmt.Println(param_2)
	obj.AddNum(-3)
	param_2 = obj.FindMedian()
	fmt.Println(param_2)
}

type MedianFinder struct {
	queMin, queMax hp
}

func ConstructorMedianFinder() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
	minQ, maxQ := &mf.queMin, &mf.queMax
	if minQ.Len() == 0 || num <= -minQ.IntSlice[0] {
		heap.Push(minQ, -num)
		if maxQ.Len()+1 < minQ.Len() {
			heap.Push(maxQ, -heap.Pop(minQ).(int))
		}
	} else {
		heap.Push(maxQ, num)
		if maxQ.Len() > minQ.Len() {
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	minQ, maxQ := mf.queMin, mf.queMax
	if minQ.Len() > maxQ.Len() {
		return float64(-minQ.IntSlice[0])
	}
	return float64(maxQ.IntSlice[0]-minQ.IntSlice[0]) / 2
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

//type MedianFinder struct {
//	nums []int
//}
//
///** initialize your data structure here. */
//func ConstructorMedianFinder() MedianFinder {
//	return MedianFinder{
//		nums: make([]int, 0),
//	}
//}
//
//func (this *MedianFinder) AddNum(num int) {
//	n := len(this.nums)
//	index := sort.SearchInts(this.nums, num)
//	left, right := make([]int, index), make([]int, n-index)
//	copy(left, this.nums[:index])
//	copy(right, this.nums[index:])
//	this.nums = append(append(left, num), right...)
//}
//
//func (this *MedianFinder) FindMedian() float64 {
//	n := len(this.nums)
//	if n%2 == 0 {
//		return (float64(this.nums[n/2]) + float64(this.nums[n/2-1])) / 2
//	} else {
//		return float64(this.nums[n/2])
//	}
//}
