package Array

import (
	"testing"
)

//二分搜索
func searchRange(nums []int, target int) []int {
	//找最左
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		val := nums[mid]
		if val == target {
			right = mid - 1
		} else if val > target {
			right = mid - 1
		} else if val < target {
			left = mid + 1
		}
	}
	nums1 := left
	if left < 0 || left >= len(nums) || nums[left] != target {
		nums1 = -1
	}
	left, right = 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		val := nums[mid]
		if val == target {
			left = mid + 1
		} else if val > target {
			right = mid - 1
		} else if val < target {
			left = mid + 1
		}
	}
	nums2 := right
	if right < 0 || right >= len(nums) || nums[right] != target {
		nums2 = -1
	}

	return []int{nums1, nums2}
}

func Test_5(t *testing.T) {
}
