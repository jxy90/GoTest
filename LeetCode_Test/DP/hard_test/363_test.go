package DP_test

import (
	"fmt"
	"math"
	"testing"
)

func maxSumSubmatrix(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	res := math.MinInt32
	for i := 1; i <= m; i++ {
		for j := i; j <= m; j++ {
			for p := 1; p <= n; p++ {
				for q := p; q <= n; q++ {
					temp := dp[j][q] - dp[i-1][q] - dp[j][p-1] + dp[i-1][p-1]
					if temp <= k && temp > res {
						res = temp
					}
				}
			}
		}
	}
	return res
}

func TestName(t *testing.T) {
	//fmt.Println(maxSumSubmatrix([][]int{{1, 0, 1}, {0, -2, 3}}, 2))
	fmt.Println(maxSumSubmatrix([][]int{{2, 2, -1}}, 3))
}
