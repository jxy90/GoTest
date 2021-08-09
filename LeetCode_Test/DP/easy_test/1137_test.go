package easy_test

import (
	"testing"
)

func Test_tribonacci(t *testing.T) {
	println(tribonacci(4))
}

var f [38]int

func Init() {
	f[0], f[1], f[2] = 0, 1, 1
	for i := 3; i <= 37; i++ {
		f[i] = f[i-1] + f[i-2] + f[i-3]
	}
}

func tribonacci(n int) int {
	if f[4] == 0 {
		Init()
	}
	return f[n]
}

func tribonacci0(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}
	f := make([]int, n+1)
	f[0], f[1], f[2] = 0, 1, 1
	for i := 3; i <= n; i++ {
		f[i] = f[i-1] + f[i-2] + f[i-3]
	}
	return f[n]
}
