package easy_test

import "testing"

func Test_hasAlternatingBits(t *testing.T) {
	println(hasAlternatingBits(3))
	println(hasAlternatingBits(10))
	println(hasAlternatingBits(11))
	println(hasAlternatingBits(7))
	println(hasAlternatingBits(5))
}

func hasAlternatingBits(n int) bool {
	last := -1
	for n != 0 {
		if last == -1 {
			last = n & 1
			continue
		}
		n >>= 1
		if last == n&1 {
			return false
		}
		last = n & 1
	}
	return true
}
