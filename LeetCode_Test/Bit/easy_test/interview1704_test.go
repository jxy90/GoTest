package easy_test

import "testing"

func Test_missingNumber2(t *testing.T) {
	println(missingNumber([]int{3, 0, 1}))
}

func missingNumber2(nums []int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < n+1; i++ {
		ans ^= i
		if i != n {
			ans ^= nums[i]
		}
	}
	return ans
}
