package DP_test

import (
	"fmt"
	"math"
	"testing"
)

/*
leetcode265. 粉刷房子 II
假如有一排房子，共 n 个，每个房子可以被粉刷成 k 种颜色中的一种，你需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同。

当然，因为市场上不同颜色油漆的价格不同，所以房子粉刷成不同颜色的花费成本也是不同的。每个房子粉刷成不同颜色的花费是以一个 n*k 的矩阵来表示的。

例如，costs[0][0] 表示第 0 号房子粉刷成 0 号颜色的成本花费；costs[1][2] 表示第 1 号房子粉刷成 2 号颜色的成本花费，以此类推。请你计算出粉刷完所有房子最少的花费成本。

注意：

所有花费均为正整数。

示例：

输入: [[1,5,3],[2,9,4]]
输出: 5
解释: 将 0 号房子粉刷成 0 号颜色，1 号房子粉刷成 2 号颜色。最少花费: 1 + 4 = 5;
     或者将 0 号房子粉刷成 2 号颜色，1 号房子粉刷成 0 号颜色。最少花费: 3 + 2 = 5.
进阶：

您能否在 O(nk) 的时间复杂度下解决此问题？
*/

func Test_minCostII(t *testing.T) {
	//2 + 5 + 3 = 10
	costs := [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}
	fmt.Println(minCostII(costs))
	//1 + 4 = 5
	//3 + 2 = 5
	costs = [][]int{{1, 5, 3}, {2, 9, 4}}
	fmt.Println(minCostII(costs))
}

func minCostII(costs [][]int) int {
	n := len(costs)
	if n == 0 {
		return 0
	}
	m := len(costs[0])
	//dp初始化
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i := range dp[0] {
		dp[0][i] = 0
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < m; j++ {
			dp[i][j] = costs[i-1][j] + minCostIIHelper(dp, i-1, j)
		}
	}
	ans := math.MaxInt32
	for i := range dp[n] {
		if dp[n][i] < ans {
			ans = dp[n][i]
		}
	}
	return ans
}

func minCostIIHelper(dp [][]int, i, j int) int {
	ans := math.MaxInt32
	for index := range dp[i] {
		if index != j && dp[i][index] < ans {
			ans = dp[i][index]
		}
	}
	return ans
}
