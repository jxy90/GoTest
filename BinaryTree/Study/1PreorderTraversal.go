package main

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

func (root *TreeNode) preorderTraversal3() []int {
	var stack []*TreeNode
	var result []int
	if root == nil {
		return nil
	}
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

// 先序遍历-非递归
func (root *TreeNode) preorderTraversal2() []int {
	if root == nil {
		return nil
	}
	var rootTemp = root
	var stack []*TreeNode
	var ret []int
	for rootTemp != nil || len(stack) != 0 {
		if rootTemp != nil {
			ret = append(ret, rootTemp.Val)
			stack = append(stack, rootTemp)
			rootTemp = rootTemp.Left
		} else {
			rootTemp = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			//ret = append(ret, rootTemp.Val)
			rootTemp = rootTemp.Right
		}
	}
	return ret
}
