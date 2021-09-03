package middle_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_longestCommonSubsequence(t *testing.T) {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] {
				f[i][j] = f[i-1][j-1] + 1
			} else {
				f[i][j] = CommonUtil.Max(f[i-1][j], f[i][j-1])
			}
		}
	}
	return f[m][n]
}
