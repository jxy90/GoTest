package middle_test

import (
	"sort"
	"testing"
)

func Test_hIndex(t *testing.T) {
	fmt.Println(hIndex([]int{1}))
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
}

func hIndex(citations []int) int {
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

func hIndex1(citations []int) int {
	n := len(citations)
	count := make([]int, n+1)
	for _, citation := range citations {
		if citation >= n {
			count[n]++
		} else {
			count[citation]++
		}
	}
	for i, tot := n, 0; i >= 0; i-- {
		tot += count[i]
		if tot >= i {
			return i
		}
	}
	return 0
}

func hIndex0(citations []int) int {
	sort.Ints(citations)
	n := len(citations)
	h := 0
	index := n - 1
	for index >= 0 && citations[index] > h {
		h++
		index--
	}
	return h
}
