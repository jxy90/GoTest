package easy_test

import "testing"

func Test_insertBits(t *testing.T) {
	println(insertBits(1024, 19, 2, 6))
}

func insertBits(N int, M int, i int, j int) int {
	M <<= i
	for k := i; k <= j; k++ {
		if N&(1<<k) != 0 {
			N ^= 1 << k
		}
	}
	return N | M
}
