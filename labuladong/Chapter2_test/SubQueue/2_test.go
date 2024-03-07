package SubQueue

import (
	"fmt"
	"testing"
)

//53. 最大子数组和
//https://leetcode.cn/problems/maximum-subarray/description/
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}

	return max(dp...)
}

func Test_2(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
