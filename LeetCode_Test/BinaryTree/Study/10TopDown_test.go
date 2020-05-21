package main_test

var answers int

func TopDown(root *TreeNode, level int) int {
	if root == nil {
		return level
	}
	LeftDepth := TopDown(root.Left, level+1)
	RightDepth := TopDown(root.Right, level+1)
	if LeftDepth > RightDepth {
		return LeftDepth
	} else {
		return RightDepth
	}
}
