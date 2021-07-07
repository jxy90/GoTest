package Backpack

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

/*
有 n 个物品和一个大小为 m 的背包. 给定数组 A 表示每个物品的大小和数组 V 表示每个物品的价值.

问最多能装入背包的总价值是多大?
*/

//物品1个,01背包
func Test_backPackII(t *testing.T) {
	println(backPackII(10, []int{2, 3, 5, 7}, []int{1, 5, 2, 4}))
}

func backPackII(m int, A []int, V []int) int {
	n := len(A)
	f := make([]int, m+1)
	for i := 0; i < n; i++ {
		//从大到小,反向循环,保证一个物品只用一次
		for j := m; j >= A[i]; j-- {
			f[j] = CommonUtil.Max(f[j], f[j-A[i]]+V[i])
		}
	}
	return f[m]
}

func backPackII0(m int, A []int, V []int) int {
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
				dp[i+1][j] = CommonUtil.Max(dp[i][j], dp[i][j-A[i]]+V[i])
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}
	return dp[n][m]
}
