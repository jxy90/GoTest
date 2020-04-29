package main

//左-》右-》中
func postorderTraversal(root *TreeNode) []int {
	var temp []int
	if root == nil {
		return temp
	}
	if root.Left != nil {
		temp = append(temp, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		temp = append(temp, postorderTraversal(root.Right)...)
	}
	temp = append(temp, root.Val)
	return temp
}

func postorderTraversal2(root *TreeNode) []int {
	var temp []int
	if root == nil {
		return temp
	}
	if root.Left != nil {
		temp = append(temp, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		temp = append(temp, postorderTraversal(root.Right)...)
	}
	temp = append(temp, root.Val)
	return temp
}
