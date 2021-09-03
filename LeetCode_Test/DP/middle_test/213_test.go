package middle_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return CommonUtil.Max(nums[0], nums[1])
	}
	return CommonUtil.Max(robHelper(nums[:n-1]), robHelper(nums[1:]))
}

func robHelper(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	//dp[i]表示前i个房子最大收益
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= n; i++ {
		dp[i] = CommonUtil.Max(nums[i-1]+dp[i-2], dp[i-1])
	}
	return dp[n]
}

func Test_rob(t *testing.T) {
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
}
