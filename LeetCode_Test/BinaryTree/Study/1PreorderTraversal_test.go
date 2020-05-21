package main_test

func (bt *TreeNode) preorderTraversal() []int {
	//中-》左-》右
	var temp []int
	if bt == nil {
		return temp
	}
	temp = append(temp, bt.Val)
	if bt.Left != nil {
		temp = append(temp, bt.Left.preorderTraversal()...)
	}
	if bt.Right != nil {
		temp = append(temp, bt.Right.preorderTraversal()...)
	}
	return temp
}

//非递归
func (root *TreeNode) preorderTraversalNew() []int {
	if root == nil {
		return nil
	}
	var stack []*TreeNode
	var result []int
	temp := root
	//中-》左-》右
	for temp != nil || len(stack) != 0 {
		if temp != nil {
			result = append(result, temp.Val)
			stack = append(stack, temp)
			temp = temp.Left
		} else {
			temp = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}
	return result
}
