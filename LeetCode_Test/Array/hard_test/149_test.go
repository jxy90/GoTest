package hard_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_maxPoints(t *testing.T) {
	println(maxPoints([][]int{{1, 1}, {2, 2}, {3, 3}}))
}

func maxPoints(points [][]int) int {
	n := len(points)
	ans := 1
	for i := 0; i < n; i++ {
		x := points[i]
		for j := i + 1; j < n; j++ {
			y := points[j]
			cnt := 2
			for k := j + 1; k < n; k++ {
				z := points[k]
				s1 := (x[0] - y[0]) * (y[1] - z[1])
				s2 := (x[1] - y[1]) * (y[0] - z[0])
				if s1 == s2 {
					cnt++
				}
			}
			ans = CommonUtil.Max(ans, cnt)
		}
	}
	return ans
}
