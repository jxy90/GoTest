package DP_test

import "github.com/jxy90/GoTest/Utils/CommonUtil"

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	//f[i][j] 代表考虑 s1s1 的前 i - 1i−1 个字符、考虑 s2s2 的前 j - 1j−1 的字符，形成的最长公共子序列长度。
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			//s1[i-1]==s2[j-1] : f[i][j]=f[i-1][j-1]+1f[i][j]=f[i−1][j−1]+1。代表使用 s1[i-1]s1[i−1] 与 s2[j-1]s2[j−1]形成最长公共子序列的长度。
			//s1[i-1]!=s2[j-1] : f[i][j]=max(f[i-1][j], f[i][j-1])f[i][j]=max(f[i−1][j],f[i][j−1])。代表不使用 s1[i-1]s1[i−1] 形成最长公共子序列的长度、不使用 s2[j-1]s2[j−1] 形成最长公共子序列的长度。这两种情况中的最大值。
			if text1[i-1] == text2[j-1] {
				f[i][j] = f[i-1][j-1] + 1
			} else {
				f[i][j] = CommonUtil.Max(f[i-1][j], f[i][j-1])
			}
		}
	}
	return f[m][n]
}
