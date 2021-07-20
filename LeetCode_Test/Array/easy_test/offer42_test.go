package easy_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	n := len(nums)
	f := make([]int, n+1)
	f[0] = nums[0]
	max := nums[0]
	for i := 1; i < n; i++ {
		f[i] = CommonUtil.Max(nums[i], f[i-1]+nums[i])
		max = CommonUtil.Max(max, f[i])
	}
	return max
}

func maxSubArray0(nums []int) int {
	n := len(nums)
	sums := make([]int, n+1)
	for i, num := range nums {
		sums[i+1] = sums[i] + num
	}
	max := nums[0]
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			max = CommonUtil.Max(max, sums[j]-sums[i])
		}
	}
	return max
}
