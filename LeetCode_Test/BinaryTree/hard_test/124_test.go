package hard

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"math"
	"testing"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans = math.MinInt32

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := CommonUtil.Max(0, maxPathSum(root.Left))
	right := CommonUtil.Max(0, maxPathSum(root.Right))
	ans = CommonUtil.Max(ans, left+root.Val+right)
	return CommonUtil.Max(left, right) + root.Val
}

func Test_maxPathSum(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	print(maxPathSum2(root))
}

func maxPathSum2(root *TreeNode) int {
	ans := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := Max(0, maxGain(node.Left))
		right := Max(0, maxGain(node.Right))
		ans = Max(ans, left+node.Val+right)
		return Max(left, right) + node.Val
	}
	maxGain(root)
	return ans
}
func maxPathSum3(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftGain := Max(maxGain(node.Left), 0)
		rightGain := Max(maxGain(node.Right), 0)
		priceNewPath := node.Val + leftGain + rightGain
		maxSum = Max(maxSum, priceNewPath)
		return node.Val + Max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}

func Max(args ...int) int {
	max := args[0]
	for _, item := range args {
		if item > max {
			max = item
		}
	}
	return max
}
