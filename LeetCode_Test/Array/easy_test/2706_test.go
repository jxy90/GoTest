package easy_test

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func Test_2706(t *testing.T) {
	buyChoco([]int{1, 2, 2}, 3)
}

func buyChoco(prices []int, money int) int {
	if len(prices) < 2 {
		return money
	}
	f, s := math.MaxInt32, math.MaxInt32
	for _, price := range prices {
		if price < f {
			s = f
			f = price
		} else if price < s {
			s = price
		}
	}
	if f+s <= money {
		return money - f - s
	}
	return money
}

func buyChoco1(prices []int, money int) int {
	sort.Ints(prices)
	fmt.Println(prices)
	if len(prices) > 1 && prices[0]+prices[1] <= money {
		return money - (prices[0] + prices[1])
	}
	return money
}
