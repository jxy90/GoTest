package middle_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_90(t *testing.T) {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	track := make([]int, 0)
	var backtrack func(start int)
	backtrack = func(start int) {
		ans = append(ans, append([]int{}, track...))
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			track = append(track, nums[i])
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}

	backtrack(0)
	return ans
}
