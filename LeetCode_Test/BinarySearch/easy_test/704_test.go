package easy_test

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}
func search(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		index := (left + right) >> 1
		mid := nums[index]
		if mid == target {
			return index
		} else if mid < target {
			left = index + 1
		} else {
			right = index - 1
		}
	}
	return -1
}
