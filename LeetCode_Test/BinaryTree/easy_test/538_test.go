package easy_test

import "testing"

var lastValue int

func convertBST(root *TreeNode) *TreeNode {
	lastValue = 0
	convertBSTDFS(root)
	return root
}

func convertBSTDFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	convertBSTDFS(root.Right)
	lastValue = root.Val + lastValue
	root.Val = lastValue
	convertBSTDFS(root.Left)
	return root
}

func Test_convertBST(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 13,
		},
	}
	data := convertBST(root)
	fmt.Println(data)
}
