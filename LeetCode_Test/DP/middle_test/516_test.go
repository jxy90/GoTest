package middle_test

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_longestPalindromeSubseq(t *testing.T) {
	fmt.Println(longestPalindromeSubseq("bbbab"))
}
func longestPalindromeSubseq(s string) int {
	n := len(s)
	//f[i][j]表示i~j范围内回文的长度
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			if i == j {
				f[i][j] = 1
			}
		}
	}
	for length := 1; length <= n; length++ {
		for i := 0; i < n-length; i++ {
			j := i + length
			if s[i] == s[j] {
				f[i][j] = f[i+1][j-1] + 2
			} else {
				f[i][j] = CommonUtil.Max(f[i+1][j], f[i][j-1])
			}
		}
	}
	return f[0][n-1]
}
