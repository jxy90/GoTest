package Dynamic

import (
	"fmt"
	"math"
	"testing"
)

//931. 下降路径最小和
//https://leetcode.cn/problems/minimum-falling-path-sum/description/
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	m := len(matrix[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	for i := range dp[0] {
		dp[0][i] = matrix[0][i]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := j - 1; k <= j+1; k++ {
				if k < 0 {
					k = 0
				}
				if k >= m {
					continue
				}
				tmp := matrix[i][j] + dp[i-1][k]
				dp[i][j] = int(math.Min(float64(dp[i][j]), float64(tmp)))
			}
		}
	}
	res := math.MaxInt32
	for _, v := range dp[n-1] {
		res = int(math.Min(float64(res), float64(v)))
	}
	return res
}

func Test_3(t *testing.T) {
	fmt.Println(minFallingPathSum([][]int{{-19, 57}, {-40, -5}}))
	fmt.Println(minFallingPathSum([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))
}
