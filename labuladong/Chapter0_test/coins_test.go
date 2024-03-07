package Chapter0_test

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"math"
	"testing"
)

//1.暴力递归
func coinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	ans := math.MaxInt64
	for _, coin := range coins {
		subProblem := coinChange(coins, amount-coin)
		if subProblem < 0 {
			continue
		}
		ans = CommonUtil.Min(ans, subProblem+1)
	}
	if ans == math.MaxInt64 {
		return -1
	}
	return ans
}

//2.对备忘录的递归
func coinChangeMemo(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	ans := math.MaxInt64

	memo := map[int]int{}
	for _, coin := range coins {
		var subProblem int
		if memo[amount] != 0 {
			subProblem = memo[amount]
		} else {
			subProblem = coinChange(coins, amount-coin)
			memo[amount] = subProblem
		}
		if subProblem < 0 {
			continue
		}
		ans = CommonUtil.Min(ans, subProblem+1)
	}
	if ans == math.MaxInt64 {
		return -1
	}
	return ans
}

//3.dp数组的迭代算法
func coinChangeDP(coins []int, n int) int {
	memo := make([]int, n+1, n+1)
	for k := range memo {
		memo[k] = n + 1
	}
	memo[0] = 0
	for i := 1; i <= n; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			memo[i] = CommonUtil.Min(memo[i], 1+memo[i-coin])
		}
	}
	return memo[n]
}

func Test_coinsChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
	fmt.Println(coinChangeMemo(coins, amount))
	fmt.Println(coinChangeDP(coins, amount))
}
