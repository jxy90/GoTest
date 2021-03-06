package easy_test

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
	println(isPowerOfTwo(1))
	println(isPowerOfTwo(16))
	println(isPowerOfTwo(218))
}

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func isPowerOfTwo0(n int) bool {
	temp := 1
	for temp < n {
		temp *= 2
	}
	return temp == n
}
