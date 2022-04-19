package hard_test

import (
	"testing"
)

func Test_821(t *testing.T) {
}

func shortestToChar(s string, c byte) []int {
	ans := make([]int, len(s))
	index := -10000
	for i := range s {
		if c == s[i] {
			index = i
		}
		ans[i] = i - index
	}
	index = 20000
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			index = i
		}
		ans[i] = min(ans[i], index-i)
	}
	return ans
}
