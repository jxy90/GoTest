package middle_test

import (
	"testing"
)

func Test_hIndex(t *testing.T) {
	println(hIndex([]int{3, 0, 6, 1, 5}))
}

func hIndex(citations []int) int {
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
