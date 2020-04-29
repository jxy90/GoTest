package main

func BottomUp(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := BottomUp(root.Left)
	rightDepth := BottomUp(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
