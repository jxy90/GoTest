package hard_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_kthSmallestPrimeFraction(t *testing.T) {
	fmt.Println(kthSmallestPrimeFraction([]int{1, 2, 3, 5}, 3))
}

type pair struct {
	//分子,分母
	a, b int
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	pairs := make([]pair, 0)
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x, y := arr[i], arr[j]
			pairs = append(pairs, pair{x, y})
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].a*pairs[j].b < pairs[j].a*pairs[i].b
	})
	return []int{pairs[k-1].a, pairs[k-1].b}
}
