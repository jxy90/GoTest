package hard_test

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"math"
	"testing"
)

func Test_strangePrinter(t *testing.T) {
	fmt.Println(strangePrinter("aaabbb"))
	fmt.Println(strangePrinter("aba"))
}

func strangePrinter(s string) int {
	n := len(s)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		f[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				f[i][j] = f[i][j-1]
			} else {
				f[i][j] = math.MaxInt64
				for k := i; k < j; k++ {
					f[i][j] = CommonUtil.Min(f[i][j], f[i][k]+f[k+1][j])
				}
			}
		}
	}

	return f[0][n-1]
}

func strangePrinter0(s string) int {
	n := len(s)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		f[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				f[i][j] = f[i][j-1]
			} else {
				f[i][j] = math.MaxInt64
				for k := i; k < j; k++ {
					f[i][j] = CommonUtil.Min(f[i][j], f[i][k]+f[k+1][j])
				}
			}
		}
	}
	return f[0][n-1]
}
