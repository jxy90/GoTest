package Chapter1

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func maxCoins(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	nums = append(append([]int{1}, nums...), 1)

	// dp[l][j] = max(dp[l][k]+num[l]*nums[k]*nums[j]+dp[k][j])
	n := len(nums)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	//数组长度
	for l := 3; l <= n; l++ {
		//数组开始位置
		for i := 0; i <= n-l; i++ {
			for k := i + 1; k < i+l-1; k++ {
				left := dp[i][k]
				right := dp[k][i+l-1]
				dp[i][i+l-1] = CommonUtil.Max(dp[i][i+l-1], left+nums[i]*nums[k]*nums[i+l-1]+right)
			}
		}
	}

	return dp[0][n-1]
}

func Test_maxCoins(t *testing.T) {
	nums := []int{3, 1, 5, 8}
	fmt.Println(maxCoins(nums))
}
