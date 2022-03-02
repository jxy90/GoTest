package middle_test

import (
	"fmt"
	"testing"
)

func Test_39(t *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	track := make([]int, 0)
	sum := 0

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
			sum += candidates[i]
			track = append(track, candidates[i])
			backtrack(i)
			track = track[:len(track)-1]
			sum -= candidates[i]
		}
	}

	backtrack(0)
	return ans
}
