package easy_test

import (
	"fmt"
	"testing"
)

func TestRangeSumBST(t *testing.T) {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 7},
		},
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val:   0,
				Right: &TreeNode{Val: 18},
			},
		},
	}
	fmt.Println(rangeSumBST(root, 7, 15))
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	leftVal := rangeSumBST(root.Left, low, high)
	val := 0
	if root.Val >= low && root.Val <= high {
		val = root.Val
	}
	rightVal := rangeSumBST(root.Right, low, high)
	return leftVal + rightVal + val
}
