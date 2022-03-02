package middle_test

import (
	"fmt"
	"testing"
)

func Test_216(t *testing.T) {
	fmt.Println(combinationSum3(3, 7))
}

func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	track := make([]int, 0)
	sum := 0

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(track) == k && sum == n {
			ans = append(ans, append([]int{}, track...))
			return
		}
		if len(track) > k || sum > n {
			return
		}

		for i := start; i < 10; i++ {
			sum += i
			track = append(track, i)
			backtrack(i + 1)
			sum -= i
			track = track[:len(track)-1]
		}
	}

	backtrack(1)
	return ans
}
