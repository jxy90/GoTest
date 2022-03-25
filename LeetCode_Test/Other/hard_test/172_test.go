package hard_test

import (
	"fmt"
	"testing"
)

func Test_172(t *testing.T) {
	fmt.Println(trailingZeroes(0))
	fmt.Println(trailingZeroes(1))
	fmt.Println(trailingZeroes(5))
	fmt.Println(trailingZeroes(30))
}

func trailingZeroes(n int) int {
	cnt := n / 5
	ans := cnt
	for cnt > 5 {
		cnt /= 5
		ans += cnt
	}
	return ans
}

//0 1,2,3,4
//1 5,6,7,8,9
//2 10,11,12,12,14
//3 15,16,17,18,19
//4 20~24
//6 25~29		25包含两个5
//
