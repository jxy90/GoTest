package middle_test

var depth []int

func rightSideView(root *TreeNode) []int {
	depth = make([]int, 0)
	rightSideViewHelper(root, 0)
	return depth
}

func rightSideViewHelper(root *TreeNode, dep int) {
	if root == nil {
		return
	}
	rightSideViewHelper(root.Right, dep+1)
	if dep >= len(depth) {
		depth = append(depth, root.Val)
	}
	rightSideViewHelper(root.Left, dep+1)
}
