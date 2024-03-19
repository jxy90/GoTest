package _022

import (
	"fmt"
	"math"
	"testing"
)

func Test_1(t *testing.T) {
	fmt.Println("")
}

//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//假设一个二叉搜索树具有如下特征：
//节点的左子树只包含小于当前节点的数。
//节点的右子树只包含大于当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。
//		  5
//   3         7
//1    4     2   8

func Test_CheckTree(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	fmt.Println(CheckTree(root))
}

//type TreeNode struct {
//	Val         int
//	Left, Right *TreeNode
//}

func CheckTree(root *TreeNode) bool {
	return helper(root, math.MinInt32, math.MaxInt32)
}
func helper(root *TreeNode, minVal, maxVal int) bool {
	if root == nil {
		return true
	}
	if root.Val <= minVal || root.Val >= maxVal {
		return false
	}
	return helper(root.Left, minVal, root.Val) && helper(root.Right, root.Val, maxVal)
}

//var lastNode *TreeNode
//func CheckTree(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	if !CheckTree(root.Left) {
//		return false
//	}
//	if lastNode == nil {
//		lastNode = root
//	} else if root.Val < lastNode.Val {
//		return false
//	}
//	lastNode = root
//	if !CheckTree(root.Right) {
//		return false
//	}
//	return true
//}

//func CheckTree(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	leftCheck := false
//	if root.Left == nil || (root.Left != nil && root.Val > root.Left.Val) {
//		leftCheck = CheckTree(root.Left)
//	}
//	rightCheck := false
//	if root.Right == nil || (root.Right != nil && root.Val < root.Right.Val) {
//		rightCheck = CheckTree(root.Right)
//	}
//	return leftCheck && rightCheck
//}
