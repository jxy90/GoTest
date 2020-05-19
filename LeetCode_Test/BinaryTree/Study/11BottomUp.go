package main

func BottomUp(root *TreeNode) int {
	if root == nil {
		return 0
	}
	LeftDepth := BottomUp(root.Left)
	RightDepth := BottomUp(root.Right)
	if LeftDepth > RightDepth {
		return LeftDepth + 1
	} else {
		return RightDepth + 1
	}
}
