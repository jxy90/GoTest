package easy_test

import (
	"fmt"
	"testing"
)

func twoSum167(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		}
	}
	return []int{-1, -1}
}

func Test_twoSum167(t *testing.T) {
	nums := []int{1, 2, 3}
	target := 3
	fmt.Println(twoSum167(nums, target))
}
