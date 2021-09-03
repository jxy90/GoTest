package DP_test

import (
	"math"
	"testing"
)

func Test_minCost(t *testing.T) {
	//[0,0,0,0,0]
	//[[1,10],[10,1],[10,1],[1,10],[5,1]]
	//5
	//2
	//3
	houses := []int{0, 0, 0, 0, 0}
	costs := [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}
	//fmt.Println(minCost(houses, costs, 5, 2, 3))
	//[0,0,0,1]
	//[[1,5],[4,1],[1,3],[4,4]]
	//4
	//2
	//4
	houses = []int{0, 0, 0, 1}
	costs = [][]int{{1, 5}, {4, 1}, {1, 3}, {4, 4}}
	fmt.Println(minCost(houses, costs, 4, 2, 4))
}

//dp[i][j][k] 第i个房子涂j颜色 k个街区
func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	//dp[i][j][k] 第i个房子涂j颜色 k个街区
	dp := make([][][]int, m+1)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, target+1)
			dp[i][j][0] = math.MaxInt32
		}
	}
	for i := 1; i <= m; i++ {
		color := houses[i-1]
		for j := 1; j <= n; j++ {
			for k := 1; k <= target; k++ {
				//分区的数量大于房子数量,无效
				if k > i {
					dp[i][j][k] = math.MaxInt32
					continue
				}

				if color == 0 {
					//第i个房子未上色
					c := cost[i-1][j-1]
					temp1 := math.MaxInt32
					//第i间是新的分区
					for p := 1; p <= n; p++ {
						if p != j {
							temp1 = min(temp1, dp[i-1][p][k-1])
						}
					}
					//第i间方式不是新分区
					temp2 := dp[i-1][j][k]
					dp[i][j][k] = min(temp1, temp2) + c
				} else {
					//第i个房子已上色
					if color == j { // 只有与「本来的颜色」相同的状态才允许被转移
						temp1 := math.MaxInt32
						//第i间是新的分区
						for p := 1; p <= n; p++ {
							if p != j {
								temp1 = min(temp1, dp[i-1][p][k-1])
							}
						}
						//第i间方式不是新分区
						temp2 := dp[i-1][j][k]
						dp[i][j][k] = min(temp1, temp2)
					} else {
						dp[i][j][k] = math.MaxInt32
					}
				}

			}
		}
	}
	ans := math.MaxInt32
	for i := 1; i <= n; i++ {
		ans = min(ans, dp[m][i][target])
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func min(args ...int) int {
	min := args[0]
	for _, item := range args {
		if item < min {
			min = item
		}
	}
	return min
}
