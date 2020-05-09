package easy_test

import "testing"

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < L { //太小 修剪 ，包括左子树全部砍掉
		return trimBST(root.Right, L, R)
	} else if root.Val > R { //太大 修剪，包括右子树全部砍掉
		return trimBST(root.Left, L, R)
	}
	root.Left = trimBST(root.Left, L, R)   //没问题的修剪左子树
	root.Right = trimBST(root.Right, L, R) //没问题的修剪右子树
	return root
}

func Test_trimBST(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	L := 1
	R := 2
	root = trimBST(root, L, R)
	println(root)
}
