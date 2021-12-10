package hard_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_threeSum(t *testing.T) {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) (ans [][]int) {
	n := len(nums)
	if n < 3 {
		return
	}
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		//数字相同跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		left, right := i+1, n-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left > i+1 && nums[left] == nums[left-1] {
					left++
				}
				for right < n-1 && nums[right] == nums[right+1] {
					right--
				}
				right--
				left++
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return
}
