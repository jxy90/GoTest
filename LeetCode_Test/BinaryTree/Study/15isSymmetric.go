package main

//对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}

func compare(Left, Right *TreeNode) bool {
	if Left == nil {
		return Right == nil
	}
	if Right == nil {
		return false
	}
	if Left.Val != Right.Val {
		return false
	}
	return compare(Left.Left, Right.Right) && compare(Left.Right, Right.Left)
}
