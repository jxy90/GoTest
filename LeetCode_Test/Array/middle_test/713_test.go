package middle_test

import (
	"testing"
)

func Test_713(t *testing.T) {
	t.Log(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
	t.Log(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	n := len(nums)
	ans := 0
	for i := range nums {
		if nums[i] < k {
			j := i
			cj := 1
			for j < n && cj*nums[j] < k {
				cj *= nums[j]
				j++
			}
			ans += j - i
			//fmt.Println(i, j)
		}
	}
	return ans
}
