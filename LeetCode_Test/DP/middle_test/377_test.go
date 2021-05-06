package middle_test

import (
	"testing"
)

func Test_combinationSum4(t *testing.T) {
	println(largestDivisibleSubset([]int{2, 4}))
	println(largestDivisibleSubset([]int{1, 2, 3}))
}

//
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	//for i := range dp {
	//	dp[i] = make([]int, target+1)
	//}
	//长度为i，目标值为j的方案数量
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i-nums[j] >= 0 {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}
