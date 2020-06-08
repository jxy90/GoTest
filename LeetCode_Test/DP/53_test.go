package DP_test

import "testing"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := map[int]int{}
	for i := 0; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}
	res := nums[0]
	for _, v := range dp {
		res = max(res, v)
	}
	return res
}

func Test_maxSubArray(t *testing.T) {
	nums := []int{-2, 1}
	println(maxSubArray(nums))
}
