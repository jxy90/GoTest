package main

var result [][]int

func levelOrder(root *TreeNode) [][]int {
	result = make([][]int, 0)
	if root == nil {
		return result
	}
	dfsHelper(root, 0)
	return result
}

func dfsHelper(node *TreeNode, level int) {

	if len(result) <= level {
		result = append(result, make([]int, 0))
	}
	result[level] = append(result[level], node.Val)
	if node.Left != nil {
		dfsHelper(node.Left, level+1)
	}
	if node.Right != nil {
		dfsHelper(node.Right, level+1)
	}
}

var result10 [][]int

func levelLpoop(root *TreeNode, level int) [][]int {
	if root == nil {
		return nil
	}
	if result10 == nil {
		result10 = make([][]int, 0)
	}
	if level > len(result10)-1 {
		result10 = append(result10, make([]int, 0))
	}
	result10[level] = append(result10[level], root.Val)
	if root.Left != nil {
		levelLpoop(root.Left, level+1)
	}
	if root.Right != nil {
		levelLpoop(root.Right, level+1)
	}
	return result10
}
