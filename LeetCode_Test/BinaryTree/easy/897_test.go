package easy_test

import "testing"

var ans897 *TreeNode
var point897 *TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	ans897 = &TreeNode{
		Val: 0,
	}
	point897 = ans897
	increasingBSTDFS(root)
	return ans897.Right
}
func increasingBSTDFS(root *TreeNode) {
	if root == nil {
		return
	}
	increasingBSTDFS(root.Left)
	increasingBSTVisit(root)
	increasingBSTDFS(root.Right)
	return
}
func increasingBSTVisit(root *TreeNode) {
	point897.Right = &TreeNode{
		Val: root.Val,
	}
	point897 = point897.Right
}

func Test_increasingBST(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	data := increasingBST(root)
	print(data)
}
