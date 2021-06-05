package middle_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_findMaxLength(t *testing.T) {
	println(findMaxLength([]int{0, 1}))
	println(findMaxLength([]int{0, 1, 0}))
}

func findMaxLength(nums []int) int {
	f := make(map[int]int)
	f[0] = -1
	count := 0
	ans := 0
	for i, num := range nums {
		if num == 0 {
			count--
		} else {
			count++
		}
		if _, ok := f[count]; ok {
			ans = CommonUtil.Max(ans, i-f[count])
		} else {
			f[count] = i
		}
	}

	return ans
}
