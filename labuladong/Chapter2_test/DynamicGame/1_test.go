package BackPack

import (
	"testing"
)

//func calculateMinimumHP(grid [][]int) int {
//	m := len(grid)
//	n := len(grid[0])
//	dp := make([][]int, m)
//	for i := range dp {
//		dp[i] = make([]int, n)
//	}
//	dp[0][0] = grid[0][0]
//	for i := 1; i < m; i++ {
//		dp[i][0] = min(dp[i-1][0], dp[i-1][0]+grid[i][0])
//	}
//	for i := 1; i < n; i++ {
//		dp[0][i] = min(dp[0][i-1], dp[0][i-1]+grid[0][i])
//	}
//	for i := 1; i < m; i++ {
//		for j := 1; j < n; j++ {
//			if grid[i][j] >= 0 {
//				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
//			} else {
//				dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
//			}
//		}
//	}
//	return dp[m-1][n-1]
//}

func Test_1(t *testing.T) {
	//fmt.Println(calculateMinimumHP([][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}))
}