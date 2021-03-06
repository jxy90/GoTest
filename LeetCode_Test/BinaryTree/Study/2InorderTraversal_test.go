package main_test

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
		temp = stack[len(stack)-1]
		result = append(result, temp.Val)
		stack = stack[:len(stack)-1]
		temp = temp.Right
	}
	return result
}

func inorderTraversalMorris(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	var current *TreeNode
	for root != nil {
		current = root.Left
		if current != nil {
			for current.Right != nil && current.Right != root {
				current = current.Right
			}
			if current.Right == nil {
				current.Right = root
				root = root.Left
				continue
			} else {
				current.Right = nil
			}
		}
		result = append(result, root.Val)
		root = root.Right
	}
	return result
}
