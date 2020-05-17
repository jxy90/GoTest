package middle_test

var depth []int

func RightSideView(root *TreeNode) []int {
	depth = make([]int, 0)
	RightSideViewHelper(root, 0)
	return depth
}

func RightSideViewHelper(root *TreeNode, dep int) {
	if root == nil {
		return
	}
	RightSideViewHelper(root.Right, dep+1)
	if dep >= len(depth) {
		depth = append(depth, root.Val)
	}
	RightSideViewHelper(root.Left, dep+1)
}
