package easy_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_thirdMax(t *testing.T) {
	//fmt.Println(thirdMax([]int{1, 3, 5}))
	//fmt.Println(thirdMax([]int{3, 2, 1}))
	//fmt.Println(thirdMax([]int{1, 2}))
	//fmt.Println(thirdMax([]int{2, 2, 3, 1}))
	fmt.Println(thirdMax([]int{3, 2, 1}))
}

func thirdMax(nums []int) int {
	sort.Ints(nums)
	//fmt.Println(nums)
	diff := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] != nums[i+1] {
			diff++
			if diff == 3 {
				return nums[i]
			}
		}
	}
	return nums[len(nums)-1]
}
func thirdMax0(nums []int) int {
	k := 2
	n := len(nums)
	if n < k+1 {
		k = 0
	}
	findTopK(nums, 0, len(nums)-1, k)
	return nums[k]
}

func findTopK(nums []int, start, end, k int) {
	l, r := start, end
	mid := nums[(l+r)/2]
	for l <= r {
		for nums[l] > mid {
			l++
		}
		for nums[r] < mid {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	if k <= r {
		findTopK(nums, start, r, k)
	}
	if k >= l {
		findTopK(nums, l, end, k)
	}
}
