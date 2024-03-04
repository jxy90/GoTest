package Tree

import (
	"testing"
)

//验证二叉搜索树
//https://leetcode.cn/problems/validate-binary-search-tree/description/
func isValidBST(root *TreeNode) bool {
	var helper func(root, min, max *TreeNode) bool
	helper = func(root, min, max *TreeNode) bool {
		if root == nil {
			return true
		}

		if min != nil && min.Val >= root.Val {
			return false
		}
		if max != nil && max.Val <= root.Val {
			return false
		}

		return helper(root.Left, min, root) && helper(root.Right, root, max)
	}
	return helper(root, nil, nil)
}

//二叉搜索树中的搜索
//https://leetcode.cn/problems/search-in-a-binary-search-tree/description/
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	} else if root.Val == val {
		return root
	} else if root.Val > val {
		return searchBST(root.Left, val)
	} else if root.Val < val {
		return searchBST(root.Right, val)
	}
	return nil
}

//二叉搜索树中的插入操作
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	} else if root.Val > val {
		root.Left = insertIntoBST(root.Right, val)
	}
	return root
}

//删除二叉搜索树中的节点
//https://leetcode.cn/problems/delete-node-in-a-bst/description/
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		minNode := getMin(root.Right)
		root.Right = deleteNode(root.Right, minNode.Val)
		minNode.Left = root.Left
		minNode.Right = root.Right
		root = minNode
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
func getMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

//不同的二叉搜索树
//https://leetcode.cn/problems/unique-binary-search-trees/description/
func numTrees(n int) int {
	m := n + 1
	G := make([]int, m)
	G[0] = 1
	G[1] = 1
	for i := 2; i < m; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	var helper func(start, end int) []*TreeNode
	helper = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}
		trees := make([]*TreeNode, 0)
		for i := start; i <= end; i++ {
			leftTrees := helper(start, i-1)
			rightTrees := helper(i+1, end)
			for _, leftTree := range leftTrees {
				for _, rightTree := range rightTrees {
					curTree := &TreeNode{
						Val:   i,
						Left:  leftTree,
						Right: rightTree,
					}
					trees = append(trees, curTree)
				}
			}
		}
		return trees
	}
	return helper(1, n)
}

func Test_8(t *testing.T) {
	generateTrees(3)

	//
	//root := &TreeNode{
	//	Val: 5,
	//	Left: &TreeNode{
	//		Val: 3,
	//		Right: &TreeNode{
	//			Val: 4,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val: 7,
	//		Left: &TreeNode{
	//			Val: 6,
	//		},
	//		Right: &TreeNode{
	//			Val: 8,
	//		},
	//	},
	//}
	//deleteNode(root, 5)
}
