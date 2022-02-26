package easy_test

import (
	"testing"
)

func Test_maximumDifference(t *testing.T) {
	//fmt.Println(kIncreasing([]int{4, 1, 5, 2, 6, 2}, 3))
}

func maximumDifference(nums []int) int {
	n := len(nums)
	min := nums[0]
	ans := -1
	for i := 0; i < n; i++ {
		temp := nums[i] - min
		if temp > 0 && temp > ans {
			ans = temp
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	return ans
}

func maximumDifference0(nums []int) int {
	n := len(nums)
	ans := -1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			temp := nums[j] - nums[i]
			if temp > 0 && temp > ans {
				ans = temp
			}
		}
	}
	return ans
}
