package offerII_test

import (
	"sort"
	"testing"
)

func Test_84(t *testing.T) {
}

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	var dfs func(count int)
	var temp []int
	visited := make(map[int]bool)
	dfs = func(count int) {
		if count == len(nums) {
			ans = append(ans, append([]int{}, temp...))
		}
		for i := 0; i < len(nums); i++ {
			if visited[i] == true || (i > 0 && visited[i-1] == false && nums[i] == nums[i-1]) {
				continue
			}
			visited[i] = true
			temp = append(temp, nums[i])
			dfs(count + 1)
			temp = temp[:len(temp)-1]
			visited[i] = false
		}
	}
	dfs(0)
	return ans
}
