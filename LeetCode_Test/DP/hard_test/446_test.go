package DP_test

import "testing"

func Test_numberOfArithmeticSlices(t *testing.T) {
	//println(maxSumSubmatrix([][]int{{1, 0, 1}, {0, -2, 3}}, 2))
	println(numberOfArithmeticSlices([]int{2, 4, 6, 8, 10}))
}

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	f := make([]map[int]int, n)
	ans := 0
	for i, x := range nums {
		f[i] = map[int]int{}
		for j, y := range nums[:i] {
			diff := y - x
			cnt := f[j][diff]
			ans += cnt
			f[i][diff] += cnt + 1
		}
	}
	return ans
}
