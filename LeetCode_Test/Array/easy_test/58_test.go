package easy_test

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	arr := strings.Split(s, " ")
	n := len(arr)
	return len(arr[n-1])
}

func lengthOfLastWord0(s string) int {
	s = strings.TrimSpace(s)
	n := len(s)
	cnt := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		cnt++
	}
	return cnt
}
