package Chapter0_test

import (
	"testing"
)

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	permuteBackTrack(nums, make([]int, 0))
	return res
}

var res [][]int

func permuteBackTrack(nums, track []int) {
	if len(nums) == len(track) {
		temp := make([]int, len(nums))
		copy(temp, track)
		res = append(res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !permuteValidate(track, nums[i]) {
			continue
		}
		track = append(track, nums[i])
		permuteBackTrack(nums, track)
		track = track[:len(track)-1]
	}
}

func permuteValidate(track []int, num int) bool {
	for _, v := range track {
		if v == num {
			return false
		}
	}
	return true
}

func Test_permute(*testing.T) {
	nums := []int{5, 4, 6, 2}
	data := permute(nums)
	fmt.Println(data)
}
