package easy_test

import "strconv"

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameTreeDFS(p) == isSameTreeDFS(q)
}

func isSameTreeDFS(root *TreeNode) string {
	if root == nil {
		return ",nil"
	}
	isSameTreeDFS(root.Left)
	isSameTreeDFS(root.Right)
	return "," + strconv.Itoa(root.Val) + isSameTreeDFS(root.Left) + isSameTreeDFS(root.Right)
}
