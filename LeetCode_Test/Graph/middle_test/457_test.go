package middle_test_test

import (
	"fmt"
	"testing"
)

func Test_circularArrayLoop(t *testing.T) {
	fmt.Println(circularArrayLoop([]int{2, -1, 1, 2, 2}))
}
func circularArrayLoop(nums []int) bool {
	n := len(nums)
	check := func(start int) bool {
		cur := start
		flag := nums[start] > 0
		k := 1
		for {
			if k > n {
				return false
			}
			next := ((cur+nums[cur])%n + n) % n
			if flag && nums[next] < 0 {
				return false
			}
			if !flag && nums[next] > 0 {
				return false
			}
			if next == start {
				return k > 1
			}
			cur = next
			k++
		}
	}
	for i := range nums {
		if check(i) {
			return true
		}
	}
	return false
}
