package easy_test

func isCousins(root *TreeNode, x int, y int) bool {
	depth = make(map[int]int)
	parent = make(map[int]int)
	isCousinsHelper(root, 0, -1)
	return depth[x] == depth[y] && parent[x] != parent[y]
}

var depth map[int]int
var parent map[int]int

//记录每个节点的深度和父节点
func isCousinsHelper(root *TreeNode, dep, par int) {
	if root == nil {
		return
	}
	isCousinsHelper(root.Left, dep+1, root.Val)
	isCousinsHelper(root.Right, dep+1, root.Val)
	depth[root.Val] = dep
	parent[root.Val] = par
	return
}
