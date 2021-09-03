package easy_test

import "testing"

func Test_search(t *testing.T) {
	fmt.Println(search([]int{1}, 1))
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 6))
}

func search(nums []int, target int) int {
	start, end := 0, len(nums)
	if end == 0 {
		return 0
	}
	l, r := start, end-1
	//找左边界
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	left := -1
	if 0 <= r && r < end && nums[r] == target {
		left = r
	} else if 0 <= l && l < end && nums[l] == target {
		left = l
	}
	if left == -1 {
		return 0
	}
	//找右边界
	l, r = start, end-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	right := -1
	if 0 <= l && l < end && nums[l] == target {
		right = l
	} else if 0 <= r && r < end && nums[r] == target {
		right = r
	}
	if right == -1 {
		return 0
	}
	return right - left + 1
}

func search0(nums []int, target int) int {
	ans := 0
	for _, num := range nums {
		if num == target {
			ans++
		}
	}
	return ans
}
