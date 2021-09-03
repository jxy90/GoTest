package easy_test

import (
	"testing"
)

func Test_search(t *testing.T) {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}
func search(nums []int, target int) int {
	l, r := 0, len(nums)
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
