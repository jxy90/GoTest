package easy_test

import (
	"testing"
)

func Test_806(t *testing.T) {
	//fmt.Println(kIncreasing([]int{4, 1, 5, 2, 6, 2}, 3))
}

func numberOfLines(widths []int, s string) []int {
	max := 100
	cur := 0
	cnt := 1
	for i := range s {
		index := s[i] - 'a'
		width := widths[index]
		if cur+width <= max {
			cur += width
		} else {
			cur = width
			cnt++
		}
	}
	return []int{cnt, cur}
}
