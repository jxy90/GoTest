package easy_test

import "testing"

func Test_hasAlternatingBits(t *testing.T) {
	fmt.Println(hasAlternatingBits(3))
	fmt.Println(hasAlternatingBits(10))
	fmt.Println(hasAlternatingBits(11))
	fmt.Println(hasAlternatingBits(7))
	fmt.Println(hasAlternatingBits(5))
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
