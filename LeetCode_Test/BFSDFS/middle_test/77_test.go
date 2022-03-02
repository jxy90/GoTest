package middle_test

import (
	"fmt"
	"testing"
)

func Test_77(t *testing.T) {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	track := make([]int, 0)

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(track) == k {
			ans = append(ans, append([]int{}, track...))
		}
		for i := start; i <= n; i++ {
			track = append(track, i)
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}

	backtrack(1)
	return ans
}
