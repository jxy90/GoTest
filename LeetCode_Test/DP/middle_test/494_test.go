package middle_test

import (
	"testing"
)

func Test_findTargetSumWays(t *testing.T) {
	println(findTargetSumWays([]int{1}, 2))
	println(findTargetSumWays([]int{1}, 1))
	println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

func findTargetSumWays(nums []int, target int) int {
	//	//f[i][j] 前i个数和为j的 的方式数量
	s := 0
	for i := range nums {
		s += nums[i]
	}
	if target > s {
		return 0
	}
	length := len(nums)
	f := make([][]int, length+1)
	for i := range f {
		f[i] = make([]int, 2*s+1)
	}
	f[0][0+s] = 1
	for i := 1; i <= length; i++ {
		x := nums[i-1]
		for j := -s; j <= s; j++ {
			if s+j-x >= 0 {
				f[i][s+j] += f[i-1][s+j-x]
			}
			if s+j+x <= 2*s {
				f[i][s+j] += f[i-1][s+j+x]
			}
		}
	}
	return f[length][target+s]
}
