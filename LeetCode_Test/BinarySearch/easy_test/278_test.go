package easy_test

import (
	"fmt"
	"testing"
)

func Test_firstBadVersion(t *testing.T) {
	data := firstBadVersion(3)
	fmt.Println(data)
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
