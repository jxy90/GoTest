package middle_test

import (
	"fmt"
	"testing"
)

func Test_restoreArray(t *testing.T) {
	fmt.Println(restoreArray([][]int{{1, 2}, {3, 4}, {3, 2}}))
}

func restoreArray(adjacentPairs [][]int) []int {
	cache := map[int][]int{}
	for _, pair := range adjacentPairs {
		cache[pair[0]] = append(cache[pair[0]], pair[1])
		cache[pair[1]] = append(cache[pair[1]], pair[0])
	}
	n := len(adjacentPairs) + 1
	ans := make([]int, n)
	for key, val := range cache {
		if len(val) == 1 {
			ans[0] = key
			break
		}
	}
	ans[1] = cache[ans[0]][0]
	for i := 1; i < n-1; i++ {
		nexts := cache[ans[i]]
		if nexts[0] == ans[i-1] {
			ans[i+1] = nexts[1]
		} else {
			ans[i+1] = nexts[0]
		}
	}
	return ans
}
