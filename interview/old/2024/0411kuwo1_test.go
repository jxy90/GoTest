package _024

import (
	"fmt"
	"testing"
)

func Test_BinarySearch(t *testing.T) {
	nums := []int{4, 5, 6, 7, 7, 7, 9}
	fmt.Println(binarySearch(nums, 7))
}

//有序数组，二分查找，最左端的值

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (right + left) / 2
		val := nums[mid]
		if target > val {
			right = mid - 1
		} else if target < val {
			left = mid + 1
		} else if target == val && mid > 0 && target == nums[mid-1] {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
