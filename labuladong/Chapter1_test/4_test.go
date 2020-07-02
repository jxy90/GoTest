package Chapter1_test

import (
	"testing"
)

// 输入一个集合，返回是否能够分割成和相等的两个子集
func canPartition(nums []int) bool {
	totalNum := 0
	for index := range nums {
		totalNum += nums[index]
	}
	if totalNum%2 != 0 {
		return false
	}
	numCount := len(nums)

	row := numCount
	col := totalNum / 2
	dp := make([][]bool, row+1)
	for i := 0; i <= row; i++ {
		dp[i] = make([]bool, col+1)
	}
	for i := 0; i <= row; i++ {
		dp[i][0] = true
	}
	for i := 1; i <= row; i++ {
		for w := 1; w <= col; w++ {
			if w-nums[i-1] < 0 {
				dp[i][w] = dp[i-1][w]
			} else {
				dp[i][w] = dp[i-1][w] || dp[i-1][w-nums[i-1]]
			}

		}
	}
	return dp[row][col]
}

///进行状态压缩
//dp[i][j] 都是通过上一行 dp[i-1][..] 转移过来的
func canPartition2(nums []int) bool {
	totalNum := 0
	for index := range nums {
		totalNum += nums[index]
	}
	if totalNum%2 != 0 {
		return false
	}
	numCount := len(nums)

	row := numCount
	col := totalNum / 2
	dp := make([]bool, col+1)
	dp[0] = true
	for i := 1; i <= row; i++ {
		for w := col; w > 0; w-- {
			if w-nums[i-1] >= 0 {
				dp[w] = dp[w] || dp[w-nums[i-1]]
			}
		}
	}
	return dp[col]
}

func Test_canPartition(t *testing.T) {
	//dp[i][j]代表 前i个数中集合结果等于j的情况是true/false。
	nums := []int{1, 5, 11, 5}
	data := canPartition2(nums)
	println(data)
	nums = []int{1, 2, 3, 5}
	data = canPartition2(nums)
	println(data)
}
