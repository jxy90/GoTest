package middle_test

import (
	"math"
	"testing"
)

func isValidBST1(root *TreeNode) bool {
	return helper(root, math.MinInt32, math.MaxInt32)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

func Test_isValidBST1(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
	}
	isValidBST1(root)
}
