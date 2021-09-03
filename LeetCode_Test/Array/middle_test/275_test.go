package middle_test

import (
	"testing"
)

func Test_hIndexII(t *testing.T) {
	fmt.Println(hIndexII([]int{0, 1, 3, 5, 6}))
}

func hIndexII(citations []int) int {
	n := len(citations)
	left, right := 0, n-1
	validate := func(mid int) bool {
		ans := 0
		for _, citation := range citations {
			if citation >= mid {
				ans++
			}
		}
		return ans >= mid
	}
	for left <= right {
		mid := (left + right) >> 1
		if validate(mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if validate(left) {
		return left
	}
	return right
}
func hIndexII1(citations []int) int {
	n := len(citations)
	count := make([]int, n+1)
	for _, citation := range citations {
		if citation >= n {
			count[n]++
		} else {
			count[citation]++
		}
	}
	total := 0
	for i := n; i > 0; i-- {
		total += count[i]
		if total >= i {
			return i
		}
	}
	return 0
}

func hIndexII0(citations []int) int {
	n := len(citations)
	index := n - 1
	h := 0
	for index >= 0 && citations[index] >= h {
		h++
		index--
	}
	return h
}
