package BinarySearch

//二分查找
//https://labuladong.online/algo/essential-technique/binary-search-framework/

//一、寻找一个数（基本的二分搜索）
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

//二、寻找左侧边界的二分搜索
func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left < 0 || left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

//三、寻找右侧边界的二分查找
func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if right < 0 || right >= len(nums) || nums[right] != target {
		return -1
	}
	return right
}

//四、逻辑统一
//https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	nums1 := left
	if left < 0 || left > len(nums)-1 || nums[left] != target {
		nums1 = -1
	}

	left, right = 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	nums2 := right
	if right < 0 || right > len(nums)-1 || nums[right] != target {
		nums2 = -1
	}
	return []int{nums1, nums2}
}
