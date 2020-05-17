package easy_test

import "math"

var maxDep int

func maxDepth(root *TreeNode) int {
	// maxDep104 = 0
	// topDown104(root, 1)
	// return maxDep104

	return ButtomUp(root)
}

//自顶向下
func topDown(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	topDown(root.Left, depth+1)
	topDown(root.Right, depth+1)
	if maxDep < depth {
		maxDep = depth
	}
}

//自底向上
func ButtomUp(root *TreeNode) int {
	if root == nil {
		return 0
	}
	Left := ButtomUp(root.Left) + 1
	Right := ButtomUp(root.Right) + 1
	return int(math.Max(float64(Left), float64(Right)))
}
