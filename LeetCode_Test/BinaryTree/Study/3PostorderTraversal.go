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

type TagNode struct {
	Node *TreeNode
	Tag  string
}

func postOrderNew(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TagNode{}
	result := []int{}
	currentNode := root
	currentTagNode := &TagNode{}
	for len(stack) > 0 || currentNode != nil {
		for currentNode != nil {
			currentTagNode = &TagNode{
				Node: currentNode,
				Tag:  "Left",
			}
			stack = append(stack, currentTagNode)
			currentNode = currentNode.Left
		}

		currentTagNode = stack[len(stack)-1]
		//stack = stack[:len(stack)-1]
		if currentTagNode.Tag == "Left" { // 只有左子树被访问过
			currentNode = currentTagNode.Node.Right
			currentTagNode.Tag = "Right"
			//stack = append(stack, currentTagNode)
		} else { // 否则则为Right 则代表右子树也被访问了，则可以访问本结点了
			stack = stack[:len(stack)-1]
			currentNode = currentTagNode.Node
			result = append(result, currentNode.Val)
			currentNode = nil
		}
	}
	return result
}

// 后序遍历 非递归
func PostOrderNew(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}
	stack := []*TreeNode{}
	current := root
	var pre (*TreeNode)
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		current = stack[len(stack)-1]
		if current.Right == nil || pre == current.Right {
			result = append(result, current.Val)
			pre = current
			stack = stack[:len(stack)-1]
			current = nil
		} else {
			current = current.Right
		}
	}
	return result
}

// 后序遍历2 非递归
func PostOrderNew2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	s1, s2 := []*TreeNode{}, []*TreeNode{}
	s1 = append(s1, root)
	for len(s1) > 0 {
		current := s1[len(s1)-1]
		s1 = s1[:len(s1)-1]
		s2 = append(s2, current)
		if current.Left != nil {
			s1 = append(s1, current.Left)
		}
		if current.Right != nil {
			s1 = append(s1, current.Right)
		}
	}
	result := []int{}
	for len(s2) > 0 {
		current := s2[len(s2)-1]
		s2 = s2[:len(s2)-1]
		result = append(result, current.Val)
	}
	return result
}
