package BackPack

import (
	"fmt"
	"testing"
)

//518. 零钱兑换 II
func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}

		}
	}
	return dp[n][amount]
}

func Test_2(t *testing.T) {
	fmt.Println(change(5, []int{1, 2, 5}))
}
