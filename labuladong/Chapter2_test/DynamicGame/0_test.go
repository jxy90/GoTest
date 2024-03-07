package BackPack

import (
	"fmt"
	"testing"
)

func max(nums ...int) int {
	val := nums[0]
	for _, v := range nums {
		if val < v {
			val = v
		}
	}
	return val
}
func min(nums ...int) int {
	val := nums[0]
	for _, v := range nums {
		if val > v {
			val = v
		}
	}
	return val
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < len(dp); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i][j-1]+grid[i][j], dp[i-1][j]+grid[i][j])
		}
	}
	return dp[m-1][n-1]
}

func Test_0(t *testing.T) {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
