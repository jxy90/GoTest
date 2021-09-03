package easy_test

import "testing"

func Test_findErrorNums(t *testing.T) {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
}

func findErrorNums(nums []int) []int {
	n := len(nums) + 1
	count := make([]int, n)
	for i := range nums {
		count[nums[i]]++
	}
	ans := []int{0, 0}
	for i := range count {
		if count[i] == 2 {
			ans[0] = i
		}
		if count[i] == 0 {
			ans[1] = i
		}
	}
	return ans
}
