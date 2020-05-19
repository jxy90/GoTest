package main

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	temp := sum - root.Val
	if root.Left == nil && root.Right == nil {
		return temp == 0
	}
	return hasPathSum(root.Left, temp) || hasPathSum(root.Right, temp)
}

func hasPathSum2(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	temp := sum - root.Val
	if root.Left == nil && root.Right == nil {
		return temp == 0
	}
	return hasPathSum(root.Left, temp) || hasPathSum(root.Right, temp)
}
