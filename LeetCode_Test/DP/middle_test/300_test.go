package middle_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func LISDP(nums []int) int {
	dp := map[int]int{}
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = CommonUtil.Max(dp[i], dp[j]+1)
			}
		}
	}
	maxLength := 0
	for _, v := range dp {
		maxLength = CommonUtil.Max(maxLength, v)
	}
	return maxLength
}

func Test_LIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	println(LISDP(nums))
}
