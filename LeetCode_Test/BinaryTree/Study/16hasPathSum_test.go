package main_test

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	temp := sum - root.Val
	if root.Left == nil && root.Right == nil {
		return temp == 0
	}
	return hasPathSum(root.Left, temp) || hasPathSum(root.Right, temp)
}

func maxPath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxPath(root.Left)
	right := maxPath(root.Right)

	return root.Val + CommonUtil.Max(left, right)
}

func Test_maxPath(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 8,
		},
	}
	fmt.Println(maxPath(root))
}
