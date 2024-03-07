package BackPack

import (
	"fmt"
	"testing"
)

//121. 买卖股票的最佳时机
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

//122. 买卖股票的最佳时机 II
func maxProfit2(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func maxProfitFreezing(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[1][0] = max(0)
	dp[1][1] = -prices[1]
	for i := 2; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

//买卖股票的最佳时机 III
func maxProfit3(prices []int) int {
	n := len(prices)
	dp := make([][3][2]int, n)
	dp[0][0][0] = 0
	dp[0][1][0] = 0
	dp[0][0][1] = 0
	dp[0][1][1] = -prices[0]
	for i := 1; i < n; i++ {
		for j := 1; j < 3; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j-1][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[n-1][2][0]
}

//188. 买卖股票的最佳时机 IV
func maxProfit4(k int, prices []int) int {
	//dp[i][j][k]第i天进行了j次交易，我们手上是否有股票。k（0表示没有,1表示有）
	n := len(prices)
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			for p := 0; p <= 1; p++ {
				//dp[i][j][p] = max(dp[i])
			}
		}
	}
	return 0
}

func Test_10(t *testing.T) {
	//fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	//fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	//fmt.Println(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
	//fmt.Println(maxProfit2([]int{1, 2, 3, 4, 5}))
	//fmt.Println(maxProfit2([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfitFreezing([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfitFreezing([]int{1}))
	//fmt.Println(maxProfit3([]int{1, 2, 3, 4, 5}))
	//fmt.Println(maxProfit3([]int{7, 6, 4, 3, 1}))
	//fmt.Println(maxProfit3([]int{1}))
}
