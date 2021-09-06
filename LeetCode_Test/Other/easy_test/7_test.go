package easy_test

import (
	"fmt"
	"math"
	"testing"
)

func Test_reverse(t *testing.T) {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(0))
	fmt.Println(reverse(10))
	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {
	ans := 0
	for x != 0 {
		//优化到此处
		if ans > math.MaxInt32/10 || ans < math.MinInt32/10 {
			return 0
		}
		ans = ans*10 + x%10
		x /= 10
	}

	return ans
}
func reverse1(x int) int {
	ans := 0
	for x != 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	//此处ans 可能超过int32限制
	if ans > math.MaxInt32 || ans < math.MinInt32 {
		return 0
	}
	return ans
}
