package easy_test

import (
	"testing"
)

func Test_firstBadVersion(t *testing.T) {
	data := firstBadVersion(3)
	println(data)
}

func isBadVersion(version int) bool {
	return true
}
func firstBadVersion(n int) int {
	m := 0
	for m <= n {
		temp := (n + m) / 2
		if isBadVersion(temp) {
			n = temp - 1
		} else {
			m = temp + 1
		}
	}
	return m
}
