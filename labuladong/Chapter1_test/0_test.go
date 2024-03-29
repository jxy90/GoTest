package Chapter1_test

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"math"
	"testing"
)

//动态规划解题套路框架
func fib(N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	return fib(N-1) + fib(N-2)
}

func fibHis(N int) int {
	his := map[int]int{}
	if N == 1 || N == 2 {
		return 1
	}
	if _, ok := his[N]; ok {
		return his[N]
	}
	tempN := fibHis(N-1) + fibHis(N-2)
	return tempN
}

func fibDP(N int) int {
	dp := map[int]int{}
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= N; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[N]
}

func fibDP2(N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	pre := 1
	cur := 1
	for i := 3; i <= N; i++ {
		temp := cur
		cur = pre + cur
		pre = temp
	}
	return cur
}

func Test_fib(t *testing.T) {
	fmt.Println(fib(5))
	fmt.Println(fibDP(5))
	fmt.Println(fibHis(5))
	fmt.Println(fibDP2(5))
}

func coinChange(coins []int, amount int) int {
	res := math.MaxInt32
	if amount < 0 {
		return -1
	} else if amount == 0 {
		return 0
	}
	for _, coin := range coins {
		sub := coinChange(coins, amount-coin)
		if sub < 0 {
			continue
		}
		res = CommonUtil.Min(res, sub+1)
	}

	if res == math.MaxInt32 {
		return -1
	}
	return res
}
func coinChangeMemo(coins []int, amount int) int {
	var memo map[int]int
	res := math.MaxInt32
	if amount < 0 {
		return -1
	} else if amount == 0 {
		return 0
	}
	for _, coin := range coins {
		var sub int
		if _, ok := memo[amount]; ok {
			sub = memo[amount]
		} else {
			sub = coinChange(coins, amount-coin)
		}
		if sub < 0 {
			continue
		}
		res = CommonUtil.Min(res, sub+1)
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func coinChangeDP(coins []int, amount int) int {
	dp := map[int]int{}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for _, coin := range coins {
			if amount-coin < 0 {
				continue
			}
			dp[i] = CommonUtil.Min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func Test_coinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
	fmt.Println(coinChangeMemo(coins, amount))
	fmt.Println(coinChangeDP(coins, amount))
}
