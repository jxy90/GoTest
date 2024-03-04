package Tree

import (
	"testing"
)

//寻找第 K 小的元素
//https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/
func kthSmallest(root *TreeNode, k int) int {
	ans := 0
	cnt := 0
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		cnt++
		if k == cnt {
			ans = root.Val
		}
		helper(root.Right)
	}
	helper(root)
	return ans
}

//把二叉搜索树转换为累加树
//https://leetcode.cn/problems/convert-bst-to-greater-tree/description/

func convertBST(root *TreeNode) *TreeNode {
	//右左中的顺序
	sum := 0
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Right)
		sum += root.Val
		root.Val = sum
		helper(root.Left)
	}
	helper(root)
	return root
}

func Test_7(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 6,
		},
	}
	convertBST(root)
}
