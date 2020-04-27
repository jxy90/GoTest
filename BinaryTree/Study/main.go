package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var node1, node2, node3 *TreeNode
	(node1).Val = 1
	(node2).Val = 2
	(node3).Val = 3

	//node1.Right = node2
	//node2.Left = node3

	println(preorderTraversal(node1))
}
