package easy_test

import "testing"

func Test_xorOperation(t *testing.T) {
	println(xorOperation(5, 0))
}
func xorOperation(n int, start int) int {
	ans := 0
	for i := 0; i < n; i++ {
		ans ^= start + i*2
	}
	return ans
}
