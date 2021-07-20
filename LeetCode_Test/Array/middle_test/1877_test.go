package middle_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"sort"
	"testing"
)

func Test_minPairSum(t *testing.T) {
	println(minPairSum([]int{3, 5, 2, 3}))
	println(minPairSum([]int{3, 5, 4, 2, 4, 6}))
}

func minPairSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := nums[0]
	for i := 0; i < n/2; i++ {
		ans = CommonUtil.Max(ans, nums[i]+nums[n-i-1])
	}
	return ans
}
