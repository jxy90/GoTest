package SubQueue

import (
	"fmt"
	"testing"
)

//516. 最长回文子序列
//https://leetcode.cn/problems/longest-palindromic-subsequence/description/
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			if i == j {
				dp[i][j] = 1
			}
		}
	}
	for length := 1; length < n; length++ {
		for start := 0; start < n-length; start++ {
			end := start + length
			if s[start] == s[end] {
				dp[start][end] = dp[start-1][end-1] + 2
			} else {
				dp[start][end] = max(dp[start][end-1], dp[start+1][end])
			}
		}
	}

	return dp[0][n-1]
}

func Test_4(t *testing.T) {
	fmt.Println(minimumDeleteSum("sea", "eat"))
}
