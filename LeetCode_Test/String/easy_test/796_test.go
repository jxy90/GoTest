package easy_test

import (
	"strings"
	"testing"
)

func Test_796(t *testing.T) {

}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	n := len(goal)
	for i := 0; i < n; i++ {
		x, y := goal[:i], goal[i:]
		if strings.Contains(s, x) && strings.Contains(s, y) {
			return true
		}
	}
	return false
}

func rotateString0(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}
