package main

func inorderTraversal(root *TreeNode) []int {
	//左-》中-》右
	var temp []int
	if root == nil {
		return temp
	}
	if root.Left != nil {
		temp = append(temp, inorderTraversal(root.Left)...)
	}
	temp = append(temp, root.Val)
	if root.Right != nil {
		temp = append(temp, inorderTraversal(root.Right)...)
	}
	return temp
}

func inorderTraversalNew(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	var stack []*TreeNode
	temp := root

	//左-》中-》右
	for temp != nil || len(stack) > 0 {
		for temp != nil {
			stack = append(stack, temp)
			temp = temp.Left
		}
		if len(stack) > 0 {
			temp = stack[len(stack)-1]
			result = append(result, temp.Val)
			stack = stack[:len(stack)-1]
			temp = temp.Right
		}
	}
	return result
}
