package middle_test

import "math"

var ans float64

func distributeCoins(root *TreeNode) int {
	ans = 0
	distributeCoinsDFS(root)
	return int(ans)
}

func distributeCoinsDFS(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	left := distributeCoinsDFS(root.Left)
	right := distributeCoinsDFS(root.Right)
	ans += math.Abs(left) + math.Abs(right)
	return float64(root.Val) + left + right - 1
}
