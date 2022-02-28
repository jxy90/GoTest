package middle_test

import (
	"fmt"
	"math"
	"testing"
)

func Test_divide(t *testing.T) {
	//fmt.Println(singleNumber2([]int{2, 2, 3, 2}))
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, 3))
}

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend > math.MinInt32 {
			return -dividend
		} // 只要不是最小的那个整数，都是直接返回相反数
		return math.MinInt32 // 是最小的那个，那就返回最大的整数
	}

	sign := 1
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		sign = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	res := divideHelper(dividend, divisor)
	if sign > 0 {
		if res > math.MaxInt32 {
			return math.MaxInt32
		} else {
			return res
		}
	}
	return -res
}

func divideHelper(a int, b int) int {
	if a < b {
		return 0
	}
	count := 1
	tb := b
	for tb+tb <= a {
		count = count + count
		tb = tb + tb
	}
	return count + divideHelper(a-tb, b)
}
