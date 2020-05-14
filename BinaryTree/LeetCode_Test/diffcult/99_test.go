package diffcult_test

import (
	"testing"
)

func recoverTree(root *TreeNode) {
	list = []*TreeNode{}
	recoverTreeHelper(root)
	trans := []*TreeNode{}
	for i := 0; i < len(list); i++ {
		if i == 0 && list[i].Val > list[i+1].Val {
			trans = append(trans, list[i])
		} else if i == len(list)-1 && list[i].Val < list[i-1].Val {
			trans = append(trans, list[i])
		} else if i != 0 && i != len(list)-1 && (list[i].Val < list[i-1].Val || list[i].Val > list[i+1].Val) {
			trans = append(trans, list[i])
		}
	}

	first := 0
	last := len(trans) - 1
	trans[first].Val, trans[last].Val = trans[last].Val, trans[first].Val
}

var list []*TreeNode

func recoverTreeHelper(root *TreeNode) {
	if root == nil {
		return
	}
	recoverTreeHelper(root.Left)
	list = append(list, root)
	recoverTreeHelper(root.Right)
}

func Test_recoverTree(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
			},
		},
	}
	recoverTree(root)
}
