package middle_test

import (
	"fmt"
	"math"
	"testing"
)

/*
Leetcode 256.粉刷房子
题目描述
假如有一排房子，共 n 个，每个房子可以被粉刷成红色、蓝色或者绿色这三种颜色中的一种，你需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同。

当然，因为市场上不同颜色油漆的价格不同，所以房子粉刷成不同颜色的花费成本也是不同的。每个房子粉刷成不同颜色的花费是以一个 n x 3 的矩阵来表示的。

例如，costs[0][0] 表示第 0 号房子粉刷成红色的成本花费；costs[1][2] 表示第 1 号房子粉刷成绿色的花费，以此类推。请你计算出粉刷完所有房子最少的花费成本。

注意：

所有花费均为正整数。

示例：

输入: [[17,2,17],[16,16,5],[14,3,19]]
输出: 10
解释: 将 0 号房子粉刷成蓝色，1 号房子粉刷成绿色，2 号房子粉刷成蓝色。
最少花费: 2 + 5 + 3 = 10。
*/

func Test_minCost(t *testing.T) {
	costs := [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}
	fmt.Println(minCost(costs))
}

func minCost(costs [][]int) int {
	n := len(costs)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	dp[0][0] = 0
	dp[0][1] = 0
	dp[0][2] = 0
	//dp[i][j]表示到第i个房子使用j颜色的最小耗费
	for i := 1; i <= n; i++ {
		for j := 0; j < 3; j++ {
			dp[i][j] = costs[i-1][j] + minColumn(dp, i-1, j)
		}
	}
	ans := math.MaxInt32
	for _, i := range dp[n] {
		if i < ans {
			ans = i
		}
	}
	return ans
}

func minColumn(dp [][]int, i, j int) int {
	min := math.MaxInt32
	for index := range dp[i] {
		if index != j && dp[i][index] < min {
			min = dp[i][index]
		}
	}
	return min
}
