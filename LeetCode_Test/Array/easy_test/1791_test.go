package easy_test

import (
	"fmt"
	"testing"
)

func Test_findCenter(t *testing.T) {
	fmt.Println(findCenter([][]int{{1, 2}, {2, 3}, {4, 2}}))
}

func findCenter(edges [][]int) int {
	m := 0
	cache := map[int]int{}
	for _, edge := range edges {
		for _, p := range edge {
			cache[p]++
			if p > m {
				m = p
			}
		}
	}
	for k, v := range cache {
		if v == m-1 {
			return k
		}
	}
	return 0
}
