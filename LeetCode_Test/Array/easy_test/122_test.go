package easy_test

import (
	"fmt"
	"testing"
)

func maxProfit2(prices []int) int {
	length := len(prices)
	maxProfit := 0
	for i := 0; i < length-1; i++ {
		if prices[i] < prices[i+1] {
			maxProfit += prices[i+1] - prices[i]
		}
	}
	return maxProfit
}

func Test_maxProfit2(t *testing.T) {
	array := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit2(array))
	//fmt.Println(maxProfit2New(array))
}
