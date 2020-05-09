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
	leftDepth := diameterOfBinaryTreeHelper(root.Left) + 1
	rightDepth := diameterOfBinaryTreeHelper(root.Right) + 1
	ans = int(math.Max(float64(ans), float64(leftDepth+rightDepth)))
	return int(math.Max(float64(leftDepth), float64(rightDepth)))
}
