package old

import (
	"fmt"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{1, 2, 3}))
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	visited := make([]bool, n)

	var dfs func(ints []int)
	dfs = func(ints []int) {
		if len(ints) == n {
			ans = append(ans, append([]int{}, ints...))
			return
		}
		for i := 0; i < n; i++ {
			if visited[i] || (i > 0 && !visited[i-1] && nums[i] == nums[i-1]) {
				continue
			}
			visited[i] = true
			ints = append(ints, nums[i])
			dfs(ints)
			ints = ints[:len(ints)-1]
			visited[i] = false
		}
	}
	dfs([]int{})
	return ans
}
