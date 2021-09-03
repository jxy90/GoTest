package easy_test

import (
	"testing"
)

func Test_sumBase(t *testing.T) {
	fmt.Println(sumBase(34, 6))
	//fmt.Println(sumBase(10, 10))
	//fmt.Println(sumBase(5, 10))
}

func sumBase(n int, k int) int {
	res := 0
	for n != 0 {
		res += n % k
		n /= k
	}
	return res
}
