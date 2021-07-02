package middle_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"sort"
	"testing"
)

func Test_maxIceCream(t *testing.T) {
	println(maxIceCream([]int{1, 3, 2, 4, 1}, 7))
}

func maxIceCream(costs []int, coins int) int {
	cnts := make([]int, 100001)
	ans := 0
	for _, cost := range costs {
		cnts[cost] += 1
	}
	for i := 1; i < 100001; i++ {
		if coins >= i {
			count := CommonUtil.Min(cnts[i], coins/i)
			ans += count
			coins -= count * i
		} else {
			break
		}
	}
	return ans
}

func maxIceCream0(costs []int, coins int) int {
	sort.Ints(costs)
	ans := 0
	for _, cost := range costs {
		if coins-cost >= 0 {
			coins -= cost
			ans++
		}
	}
	return ans
}
