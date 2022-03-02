package middle_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_46(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
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
