package middle_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_findUnsortedSubarray(t *testing.T) {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4}))
}

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	sortNums := make([]int, n)
	for i := range nums {
		sortNums[i] = nums[i]
	}
	sort.Ints(sortNums)
	left, right := 0, n-1
	for left < n {
		if sortNums[left] != nums[left] {
			break
		}
		left++
	}
	for right >= left {
		if sortNums[right] != nums[right] {
			break
		}
		right--
	}
	return n - left - (n - 1 - right)
}
