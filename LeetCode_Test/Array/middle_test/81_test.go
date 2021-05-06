package middle_test

import (
	"testing"
)

func Test_search(t *testing.T) {
	println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
}

func search(nums []int, target int) bool {
	start := 0
	end := len(nums) - 1
	left := start
	right := end
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		}
		if nums[left] == nums[mid] {
			left++
		} else if nums[right] == nums[mid] {
			right--
		} else if nums[left] < nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}
	return false
}
