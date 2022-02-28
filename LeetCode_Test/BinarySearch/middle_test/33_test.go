package middle_test

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	fmt.Println(search([]int{3, 1}, 1))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[0] {
			//左边有序
			if nums[0] <= target && target <= nums[mid] {
				//搜索左边
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			//右边有序
			if nums[mid] <= target && target <= nums[len(nums)-1] {
				//搜索右边
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
