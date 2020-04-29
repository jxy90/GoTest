package main

var answers int

func TopDown(root *TreeNode, level int) int {
	if root == nil {
		return level
	}
	leftDepth := TopDown(root.Left, level+1)
	rightDepth := TopDown(root.Right, level+1)
	if leftDepth > rightDepth {
		return leftDepth
	} else {
		return rightDepth
	}
}
