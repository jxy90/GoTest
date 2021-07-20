package easy_test

import (
	"testing"
)

func Test_isPerfectSquare(t *testing.T) {
	println(isPerfectSquare(9))
	println(isPerfectSquare(16))
	println(isPerfectSquare(14))
}

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	l, r := 1, num/2
	for l <= r {
		mid := l + (r-l)>>1
		mid2 := mid * mid
		if mid2 == num {
			return true
		} else if mid2 > num {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
