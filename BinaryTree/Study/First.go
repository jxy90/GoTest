package main

func preorderTraversal(root *TreeNode) []int {
	//中-》左-》右
	var temp []int
	if root == nil {
		return temp
	}
	temp = append(temp, root.Val)

	if root.Left != nil {
		temp = append(temp, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		temp = append(temp, preorderTraversal(root.Right)...)
	}
	return temp
}
