package Tree

import (
	"testing"
)

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := countNodes(root.Left)
	right := countNodes(root.Right)
	return 1 + left + right
}

func Test_10(t *testing.T) {
}
