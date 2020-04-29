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
