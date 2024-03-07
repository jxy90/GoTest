package Dynamic

import (
	"fmt"
	"math"
	"testing"
)

func fib(n int) int {
	memo := make(map[int]int)
	memo[0] = 0
	memo[1] = 1
	return fibDP(n, memo)
}

func fibDP(n int, memo map[int]int) int {
	v, ok := memo[n]
	if ok {
		return v
	}
	ans := fibDP(n-1, memo) + fibDP(n-2, memo)
	memo[n] = ans
	return ans
}

//322. 零钱兑换
//https://leetcode.cn/problems/coin-change/description/
func coinChangeMemo(coins []int, amount int) int {
	memo := make(map[int]int, 0)

	var helper func(coins []int, amount int) int
	helper = func(coins []int, amount int) int {
		res := math.MaxInt32
		if amount < 0 {
			return -1
		}
		if amount == 0 {
			return 0
		}
		if _, ok := memo[amount]; ok {
			return memo[amount]
		}

		for _, coin := range coins {
			next := helper(coins, amount-coin)
			if next < 0 {
				continue
			}
			res = int(math.Min(float64(res), float64(next+1)))
		}
		if res == math.MaxInt32 {
			memo[amount] = -1
		} else {
			memo[amount] = res
		}
		return memo[amount]
	}
	helper(coins, amount)
	return memo[amount]
}
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 0; i < amount+1; i++ {

		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coin]+1)))
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func Test_0(t *testing.T) {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}
