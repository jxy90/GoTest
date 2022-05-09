package easy_test

import (
	"testing"
)

func Test_942(t *testing.T) {
	t.Log(diStringMatch("IDID"))
}

func diStringMatch(s string) []int {
	n := len(s)
	lo, hi := 0, n
	ans := make([]int, 0, n+1)
	for i := range s {
		if s[i] == 'D' {
			ans = append(ans, hi)
			hi--
		} else {
			ans = append(ans, lo)
			lo++
		}
	}
	ans = append(ans, lo)
	return ans
}
