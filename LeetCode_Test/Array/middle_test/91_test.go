package middle_test

import (
	"github.com/jxy90/GoTest/LeetCode_Test/BinaryTree/Study/Codec"
	"math"
)

var ans int = 0

func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && ((s[i-2] == '1') || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

func maxPathSum(root *Codec.TreeNode) int {
	min := math.MinInt32
	var maxPathSumHelper func(node *Codec.TreeNode) int
	maxPathSumHelper = func(root *Codec.TreeNode) int {
		if root == nil {
			return 0
		}

		leftMax := max(maxPathSumHelper(root.Left), 0)
		rightMax := max(maxPathSumHelper(root.Right), 0)

		lrmax := max(leftMax, rightMax)

		min = max(root.Val+leftMax+rightMax, min)

		return root.Val + lrmax
	}
	return min
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
