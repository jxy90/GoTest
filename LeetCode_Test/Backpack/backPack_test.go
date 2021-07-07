package Backpack

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

/*
在n个物品中挑选若干物品装入背包，最多能装多满？假设背包的大小为m，每个物品的大小为A{i}
*/

//物品一个,01背包
func Test_backPack(t *testing.T) {
	println(backPack(10, []int{3, 4, 8, 5}))
}

func backPack(m int, A []int) int {
	n := len(A)
	f := make([]int, m+1)
	for i := 0; i < n; i++ {
		//从大到小,反向循环,保证一个物品只用一次
		for j := m; j >= A[i]; j-- {
			f[j] = CommonUtil.Max(f[j], f[j-A[i]]+A[i])
		}
	}
	return f[m]
}

func backPack0(m int, A []int) int {
	n := len(A)
	//前i个物品在大小为j的背包下,最大量
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= m; j++ {
			//使用A[i]的情况,取两个数中的最大值
			if j-A[i] >= 0 {
				dp[i+1][j] = CommonUtil.Max(dp[i][j], dp[i][j-A[i]]+A[i])
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}
	return dp[n][m]
}
