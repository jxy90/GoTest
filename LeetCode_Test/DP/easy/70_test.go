package easy_test

import (
	"testing"
)

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

//DP
func climbStairs_DP(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	pre := 1
	cur := 2
	for i := 3; i < n; i++ {
		pre, cur = cur, pre+cur
	}
	return cur
}

func Test_climbStairs(t *testing.T) {
	data := climbStairs(3)
	println(data)
}
