package easy_test

import (
	"testing"
)

func Test_numWays(t *testing.T) {
	println(numWays(5, [][]int{{0, 2}, {2, 1}, {3, 4}, {2, 3}, {1, 4}, {2, 0}, {0, 4}}, 3))
}

func numWays(n int, relation [][]int, k int) int {
	//f[i][j]表示第i轮到达第j个小朋友的方法数,返回f[k][n]
	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, n)
	}
	f[0][0] = 1
	for i := 0; i < k; i++ {
		for _, p := range relation {
			f[i+1][p[1]] += f[i][p[0]]
		}
	}
	return f[k][n-1]
}
