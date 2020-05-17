package middle_test

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	_Left, _Right = 0, 0
	all := btreeGameWinningMoveDFS(root, x)
	parent := all - _Left - _Right - 1
	if parent > n/2 || _Left > n/2 || _Right > n/2 {
		return true
	}
	return false
}

var _Left, _Right int

func btreeGameWinningMoveDFS(root *TreeNode, x int) int {
	if root == nil {
		return 0
	}
	Left := btreeGameWinningMoveDFS(root.Left, x)
	Right := btreeGameWinningMoveDFS(root.Right, x)
	if root.Val == x {
		_Left = Left
		_Right = Right
	}
	return Left + Right + 1
}
