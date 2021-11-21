package easy_test

import (
	"fmt"
	"testing"
)

func isUniValTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isUniValTreeDFS(root, root.Val)
}

func isUniValTreeDFS(root *TreeNode, Value int) bool {
	if root == nil {
		return true
	}
	if Value != root.Val {
		return false
	}
	return isUniValTreeDFS(root.Left, Value) && isUniValTreeDFS(root.Right, Value)
}

func Test_isUniValTree(t *testing.T) {
	root := &TreeNode{
		Val:  2,
		Left: nil,
		Right: &TreeNode{
			Val: 1,
		},
	}
	fmt.Println(isUniValTree(root))
}
