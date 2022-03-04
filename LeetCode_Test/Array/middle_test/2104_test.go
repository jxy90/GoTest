package middle_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_2104(t *testing.T) {
	fmt.Println(subArrayRanges([]int{1, 2, 3}))
	fmt.Println(subArrayRanges([]int{1, 3, 3}))
	fmt.Println(subArrayRanges([]int{4, -1, -3, 4, 1}))
}

func subArrayRanges(nums []int) (ans int64) {
	n := len(nums)

	for i := 0; i < n-1; i++ {
		max, min := nums[i], nums[i]
		for j := i + 1; j < n; j++ {
			if nums[j] > max {
				max = nums[j]
			}
			if nums[j] < min {
				min = nums[j]
			}
			ans += int64(max - min)
		}
	}
	return
}

//子集的情况
func subArrayRanges0(nums []int) (ans int64) {
	sort.Ints(nums)
	track := make([]int, 0)

	var backtrack func(index int)
	backtrack = func(index int) {
		if len(track) > 0 {
			fmt.Println(track)
			ans += int64(track[len(track)-1] - track[0])
		}
		for i := index; i < len(nums); i++ {
			track = append(track, nums[i])
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return
}
