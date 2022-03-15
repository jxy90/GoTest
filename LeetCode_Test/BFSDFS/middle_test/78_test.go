package middle_test

import (
	"fmt"
	"testing"
)

func Test_78(t *testing.T) {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{3, 1}))
}

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	track := make([]int, 0)
	var backtrack func(start int)
	backtrack = func(start int) {
		ans = append(ans, append([]int{}, track...))
		for i := start; i < len(nums); i++ {
			track = append(track, nums[i])
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}

	backtrack(0)
	return ans
}
