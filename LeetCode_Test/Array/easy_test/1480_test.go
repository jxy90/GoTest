package easy_test

import (
	"testing"
)

func Test_runningSum(t *testing.T) {
	println(runningSum([]int{1, 1, 0, 0, 0}))
}

func runningSum(nums []int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
