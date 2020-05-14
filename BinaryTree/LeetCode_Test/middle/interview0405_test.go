package middle_test

import "testing"

func isValidBST(root *TreeNode) bool {
	_temp = []int{}
	return isValidBSTDFS(root)

}

var _temp []int

func isValidBSTDFS(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_left := isValidBSTDFS(root.Left)
	_temp = append(_temp, root.Val)
	if len(_temp) > 1 && _temp[len(_temp)-2] >= _temp[len(_temp)-1] {
		return false
	}
	_right := isValidBSTDFS(root.Right)

	return _left && _right
}

func Test_isValidBSTDFS(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
	}
	isValidBST(root)
}
