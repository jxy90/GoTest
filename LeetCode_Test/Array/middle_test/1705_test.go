package middle_test

import (
	"container/heap"
	"fmt"
	"testing"
)

func Test_eatenApples(t *testing.T) {
	//fmt.Println(eatenApples([]int{1, 2, 3, 5, 2}, []int{3, 2, 1, 4, 2}))
	//fmt.Println(eatenApples([]int{3, 0, 0, 0, 0, 2}, []int{3, 0, 0, 0, 0, 2}))
	//fmt.Println(eatenApples([]int{2, 1, 10}, []int{2, 10, 1}))
	//fmt.Println(eatenApples([]int{9, 2}, []int{3, 5}))
	fmt.Println(eatenApples([]int{0}, []int{1}))
}

type pair struct {
	apple, day int
}

type HP []*pair

func (h *HP) Push(v interface{}) {
	*h = append(*h, v.(*pair))
}

func (h *HP) Pop() interface{} {
	length := len(*h)
	old := *h
	v := old[length-1]
	*h = old[:length-1]
	return v
}
func (h HP) Len() int {
	return len(h)
}

func (h HP) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h HP) Less(i, j int) bool {
	return h[i].day < h[j].day
}

func eatenApples(apples []int, days []int) int {
	ans := 0
	hp := &HP{}
	i := 0
	for ; i < len(apples); i++ {
		for hp.Len() > 0 && (*hp)[0].day <= i {
			heap.Pop(hp)
		}
		if apples[i] > 0 && days[i] > 0 {
			heap.Push(hp, &pair{apples[i], i + days[i]})
		}
		if hp.Len() > 0 {
			(*hp)[0].apple--
			if (*hp)[0].apple == 0 {
				heap.Pop(hp)
			}
			ans++
		}
	}
	for hp.Len() > 0 {
		for hp.Len() > 0 && (*hp)[0].day <= i {
			heap.Pop(hp)
		}
		if hp.Len() > 0 {
			(*hp)[0].apple--
			ans++
			if (*hp)[0].apple == 0 {
				heap.Pop(hp)
			}
			i++

			//p := heap.Pop(hp).(*pair)
			//num := min(p.day-i, p.apple)
			//ans += num
			//i += num
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
