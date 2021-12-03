package easy_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_largestSumAfterKNegations(t *testing.T) {
	//fmt.Println(largestSumAfterKNegations([]int{4, 2, 3}, 1))
	//fmt.Println(largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2))
	//fmt.Println(largestSumAfterKNegations([]int{-8,3,-5,-3,-5,-2}, 6))
	fmt.Println(largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2))
}

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	for i, v := range nums {
		if v >= 0 || k == 0 {
			break
		}
		nums[i] = -nums[i]
		k--
	}
	sort.Ints(nums)
	k = k % 2
	ans := 0
	for i, num := range nums {
		if i == 0 && k == 1 {
			num = -num
		}
		ans += num
	}
	return ans
}
