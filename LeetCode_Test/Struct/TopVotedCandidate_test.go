package Struct_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_TopVotedCandidate(t *testing.T) {
	//persons := []int{0, 1, 1, 0, 0, 1, 0}
	//times := []int{0, 5, 10, 15, 20, 25, 30}
	//obj := ConstructorTopVotedCandidate(persons, times)
	//fmt.Println(obj.Q(3))
	//fmt.Println(obj.Q(12))
	//fmt.Println(obj.Q(25))
	//fmt.Println(obj.Q(15))
	//fmt.Println(obj.Q(24))
	//fmt.Println(obj.Q(8))
	persons := []int{0, 0, 0, 0, 1}
	times := []int{0, 6, 39, 52, 75}
	obj := ConstructorTopVotedCandidate(persons, times)
	fmt.Println(obj.Q(45))
	fmt.Println(obj.Q(49))
	fmt.Println(obj.Q(59))
	fmt.Println(obj.Q(68))
	fmt.Println(obj.Q(42))
	fmt.Println(obj.Q(37))
	fmt.Println(obj.Q(99))
	fmt.Println(obj.Q(26))
	fmt.Println(obj.Q(78))
	fmt.Println(obj.Q(43))
}

type TopVotedCandidate struct {
	times   []int
	persons []int
	tops    []int
}

func ConstructorTopVotedCandidate(persons []int, times []int) TopVotedCandidate {
	vote := map[int]int{-1: -1}
	n := len(times)
	top := -1
	tops := make([]int, n)
	for i := 0; i < n; i++ {
		vote[persons[i]]++
		if vote[persons[i]] >= vote[top] {
			top = persons[i]
		}
		tops[i] = top
	}
	fmt.Println(times)
	fmt.Println(tops)
	return TopVotedCandidate{
		times:   times,
		persons: persons,
		tops:    tops,
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	index := sort.SearchInts(this.times, t+1) - 1
	//if index == len(this.times) || this.times[index] > t {
	//	index--
	//}
	return this.tops[index]
}
