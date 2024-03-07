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

func knapsack(W, N int, wt, val []int) int {
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
		dp[i][0] = 0
	}
	for i := 1; i <= N; i++ {
		for j := 1; j <= W; j++ {
			if j-wt[i-1] >= 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-wt[i-1]]+val[i-1])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[N][W]
}

func Test_0(t *testing.T) {
	fmt.Println(knapsack(4, 3, []int{2, 1, 3}, []int{4, 2, 3}))
}
