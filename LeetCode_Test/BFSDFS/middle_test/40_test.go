package middle_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_40(t *testing.T) {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	track := make([]int, 0)
	sum := 0
	sort.Ints(candidates)

	var backtrack func(start int)
	backtrack = func(start int) {
		if sum == target {
			ans = append(ans, append([]int{}, track...))
			return
		}
		if sum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i-1] == candidates[i] {
				continue
			}
			sum += candidates[i]
			track = append(track, candidates[i])
			backtrack(i + 1)
			track = track[:len(track)-1]
			sum -= candidates[i]
		}
	}

	backtrack(0)
	return ans
}
