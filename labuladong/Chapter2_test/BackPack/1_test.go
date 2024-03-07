package BackPack

import (
	"fmt"
	"testing"
)

//416. 分割等和子集
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	dp := make([][]bool, len(nums)+1)
	for i := range dp {
		dp[i] = make([]bool, sum/2+1)
	}
	for i := range dp {
		dp[i][0] = true
	}
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= sum/2; j++ {
			if j-nums[i-1] >= 0 {
				dp[i][j] = dp[i-1][j-nums[i-1]] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][sum/2]
}
func Test_1(t *testing.T) {
	//fmt.Println(knapsack(4, 3, []int{2, 1, 3}, []int{4, 2, 3}))
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}
