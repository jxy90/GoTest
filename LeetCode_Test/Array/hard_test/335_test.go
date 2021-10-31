package hard_test

import (
	"fmt"
	"testing"
)

func Test_isSelfCrossing(t *testing.T) {
	fmt.Println(isSelfCrossing([]int{2, 1, 1, 2}))
}

func isSelfCrossing(d []int) bool {
	n := len(d)
	if n < 4 {
		return false
	}
	for i := 3; i < n; i++ {
		if d[i] >= d[i-2] && d[i-1] <= d[i-3] {
			return true
		} else if i >= 5 && d[i-1] <= d[i-3] && d[i-2] > d[i-4] && d[i]+d[i-4] >= d[i-2] && d[i-1]+d[i-5] >= d[i-3] {
			return true
		} else if i >= 4 && d[i-1] == d[i-3] && d[i]+d[i-4] >= d[i-2] {
			return true
		}
	}
	return false
}
