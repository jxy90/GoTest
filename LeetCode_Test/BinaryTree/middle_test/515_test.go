package middle_test

import (
	"math"
	"testing"
)

func Test_515(t *testing.T) {
}

func largestValues(root *TreeNode) []int {
	ans := make([]int, 0, 50)
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(ans) == depth {
			ans = append(ans, math.MinInt32)
		}
		if ans[depth] < root.Val {
			ans[depth] = root.Val
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	return ans
}
