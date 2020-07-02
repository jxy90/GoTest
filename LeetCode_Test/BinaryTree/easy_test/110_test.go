package easy_test

import (
	"math"
	"testing"
)

func isBalanced(root *TreeNode) bool {
	//if root == nil {
	//	return true
	//}
	//if math.Abs(height(root.Left)-height(root.Right)) > 1 {
	//	return false
	//}
	//return isBalanced(root.Left) && isBalanced(root.Right)
	return isBalancedFind(root) != -1
}

func height(root *TreeNode) float64 {
	if root == nil {
		return -1
	}
	return 1 + math.Max(height(root.Left), height(root.Right))
}

func isBalancedFind(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	l := isBalancedFind(root.Left)
	r := isBalancedFind(root.Right)
	if r == -1 || l == -1 || math.Abs(l-r) > 1 {
		return -1
	}
	return math.Max(l, r) + 1
}

func Test110(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
	}
	//root = &TreeNode{
	//	Val:   1,
	//	Left:  &TreeNode{
	//		Val:   2,
	//		Left:  &TreeNode{
	//			Val:   3,
	//			Left:  &TreeNode{
	//				Val:   4,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: &TreeNode{
	//				Val:   4,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//		},
	//		Right: &TreeNode{
	//			Val:   3,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}
	println(isBalanced(root))
}
