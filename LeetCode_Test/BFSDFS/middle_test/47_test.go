package middle_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_47(t *testing.T) {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}

func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	track := make([]int, 0)
	visited := map[int]bool{}
	sort.Ints(nums)

	var backtrack func()
	backtrack = func() {
		if len(track) == len(nums) {
			ans = append(ans, append([]int{}, track...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			if i > 0 && visited[i-1] == false && nums[i] == nums[i-1] {
				continue
			}
			visited[i] = true
			track = append(track, nums[i])
			backtrack()
			visited[i] = false
			track = track[:len(track)-1]
		}
	}
	backtrack()
	return ans
}
