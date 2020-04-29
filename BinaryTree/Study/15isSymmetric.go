package main

//对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}

func compare(left, right *TreeNode) bool {
	if left == nil {
		return right == nil
	}
	if right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return compare(left.Left, right.Right) && compare(left.Right, right.Left)
}
