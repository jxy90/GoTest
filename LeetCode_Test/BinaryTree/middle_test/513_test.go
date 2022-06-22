package middle_test

import (
	"testing"
)

func Test_513(t *testing.T) {

}

func findBottomLeftValue(root *TreeNode) int {
	ans := root.Val
	level := 0
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		dfs(root.Left, depth+1)
		if depth > level {
			level = depth
			ans = root.Val
		}
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	return ans
}
