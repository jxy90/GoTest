package middle_test

import "testing"

func Test_removeDuplicates(t *testing.T) {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}

func removeDuplicates(nums []int) int {
	m := len(nums)
	ans := 0
	for i := m - 3; i >= 0; i-- {
		if nums[i] == nums[i+2] {
			ans++
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return m - ans
}
