package DP_test

import (
	"fmt"
	"testing"
)

func Test_kInversePairs(t *testing.T) {
	//fmt.Println(kInversePairs(3, 0))
	//fmt.Println(kInversePairs(3, 0))
	fmt.Println(kInversePairs(1000, 1000))
}

func kInversePairs(n int, k int) int {
	MOD := int(1e9 + 7)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, k+1)
		f[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			f[i][j] = f[i][j-1] + f[i-1][j]
			if j >= i {
				f[i][j] -= f[i-1][j-i]
			}
			if f[i][j] < 0 {
				f[i][j] += MOD
			}
			f[i][j] %= MOD
		}
	}
	return f[n][k]
}
