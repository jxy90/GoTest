package easy_test

import "math"

func diameterOfBinaryTree(root *TreeNode) int {
	ans = 0
	diameterOfBinaryTreeHelper(root)
	return ans
}

var ans int

func diameterOfBinaryTreeHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	LeftDepth := diameterOfBinaryTreeHelper(root.Left) + 1
	RightDepth := diameterOfBinaryTreeHelper(root.Right) + 1
	ans = int(math.Max(float64(ans), float64(LeftDepth+RightDepth)))
	return int(math.Max(float64(LeftDepth), float64(RightDepth)))
}
