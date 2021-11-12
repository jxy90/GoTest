package easy_test

import (
	"fmt"
	"testing"
)

func Test_maxCount(t *testing.T) {
	operations := [][]int{{2, 2}, {3, 3}}
	fmt.Println(maxCount(3, 3, operations))
}

func maxCount(m int, n int, ops [][]int) int {
	a, b := m, n
	for _, op := range ops {
		if a > op[0] {
			a = op[0]
		}
		if b > op[1] {
			b = op[1]
		}
	}
	return (a) * (b)
}
