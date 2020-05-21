package easy__test

import (
	"math"
	"testing"
)

func maxProfit(prices []int) int {
	length := len(prices)
	ans := make([]int, length)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			ans[i] = max(ans[i], prices[i]-prices[j])
		}
	}
	maxCount := 0
	for k := range ans {
		maxCount = max(maxCount, ans[k])
	}
	return maxCount
}
func maxProfitNew(prices []int) int {
	maxProfit := math.MinInt64
	minPrice := math.MaxInt64
	for i := range prices {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		maxProfit = max(maxProfit,prices[i]-minPrice)
	}

	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test_maxProfit(t *testing.T) {
	array := []int{7, 1, 5, 3, 6, 4}
	println(maxProfit(array))
	println(maxProfitNew(array))
}
