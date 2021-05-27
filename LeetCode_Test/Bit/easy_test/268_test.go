package easy_test

import "testing"

func Test_missingNumber(t *testing.T) {
	println(missingNumber([]int{1, 2, 3}))
	println(missingNumber([]int{1, 0, 3}))
	println(missingNumber([]int{0, 2, 3}))
}

func missingNumber(nums []int) int {
	n := len(nums)
	ans := n
	for i := 0; i < n; i++ {
		ans ^= i
		ans ^= nums[i]
	}
	return ans
}
