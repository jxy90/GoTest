package middle_test

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	nums := []int{-2, 1}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := map[int]int{}
	for i := 0; i < len(nums); i++ {
		dp[i] = CommonUtil.Max(nums[i], dp[i-1]+nums[i])
	}
	res := nums[0]
	for _, v := range dp {
		res = CommonUtil.Max(res, v)
	}
	return res
}
