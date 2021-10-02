package easy_test

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	nums := []int{1, 2, 3}
	target := 3
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	set := map[int]int{}
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		if _, ok := set[temp]; ok {
			return []int{set[temp], i}
		}
		set[nums[i]] = i
	}
	return []int{-1, -1}
}
