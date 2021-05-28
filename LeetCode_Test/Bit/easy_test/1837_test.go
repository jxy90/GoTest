package easy_test

import (
	"testing"
)

func Test_sumBase(t *testing.T) {
	println(sumBase(34, 6))
	//println(sumBase(10, 10))
	//println(sumBase(5, 10))
}

func sumBase(n int, k int) int {
	res := 0
	for n != 0 {
		res += n % k
		n /= k
	}
	return res
}
