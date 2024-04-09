package classic

import (
	"fmt"
	"testing"
)

func Test_findK(t *testing.T) {
	fmt.Println(findKthLargest([]int{8, 3456, 8, 34, 76, 383, 25, 6, 23, 45, 7, 8, 56, 34, 3}, 1))
}

func findKthLargest(nums []int, k int) []int {
	helper(nums, k)
	return nums
}

func helper(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	left, right := 0, len(nums)-1
	mid, i := nums[left], 1
	for left < right {
		if nums[i] > mid {
			nums[i], nums[left] = nums[left], nums[i]
			i++
			left++
		} else {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		}
	}
	helper(nums[:left], k)
	helper(nums[left+1:], k-left-1)
}
