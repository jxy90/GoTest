package easy_test

import "math"

var maxDep104 int

func maxDepth104(root *TreeNode) int {
	// maxDep104 = 0
	// topDown104(root, 1)
	// return maxDep104

	return ButtomUp104(root)
}

//自顶向下
func topDown104(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	topDown104(root.Left, depth+1)
	topDown104(root.Right, depth+1)
	if maxDep104 < depth {
		maxDep104 = depth
	}
}

//自底向上
func ButtomUp104(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := ButtomUp104(root.Left) + 1
	right := ButtomUp104(root.Right) + 1
	return int(math.Max(float64(left), float64(right)))
}
