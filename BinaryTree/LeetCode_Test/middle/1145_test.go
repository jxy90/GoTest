package middle_test

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	_left, _right = 0, 0
	all := btreeGameWinningMoveDFS(root, x)
	parent := all - _left - _right - 1
	if parent > n/2 || _left > n/2 || _right > n/2 {
		return true
	}
	return false
}

var _left, _right int

func btreeGameWinningMoveDFS(root *TreeNode, x int) int {
	if root == nil {
		return 0
	}
	left := btreeGameWinningMoveDFS(root.Left, x)
	right := btreeGameWinningMoveDFS(root.Right, x)
	if root.Val == x {
		_left = left
		_right = right
	}
	return left + right + 1
}
