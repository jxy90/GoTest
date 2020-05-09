package easy_test

import "testing"

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isUnivalTreeDFS(root, root.Val)
}

func isUnivalTreeDFS(root *TreeNode, value int) bool {
	if root == nil {
		return true
	}
	if value != root.Val {
		return false
	}
	return isUnivalTreeDFS(root.Left, value) && isUnivalTreeDFS(root.Right, value)
}

func Test_isUnivalTree(t *testing.T) {
	root := &TreeNode{
		Val:  2,
		Left: nil,
		Right: &TreeNode{
			Val: 1,
		},
	}
	println(isUnivalTree(root))
}
