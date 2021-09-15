package middle_test

import (
	"fmt"
	"testing"
)

func Test_findPeakElement(t *testing.T) {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1}))
}

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if mid-1 >= 0 && mid+1 < len(nums) && nums[mid] > nums[mid-1] && nums[mid] > nums[mid]+1 {
			return mid
		} else if mid+1 < len(nums) && nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
