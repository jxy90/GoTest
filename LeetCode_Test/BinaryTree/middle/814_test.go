package middle_test

import "testing"

func pruneTree(root *TreeNode) *TreeNode {
	//if root != nil && root.Left == nil && root.Right == nil && root.Val == 0 {
	//	return nil
	//}
	//pruneTreeDFS(root, &TreeNode{
	//	Val: 0,
	//}, "")
	//return root
	if containsOne(root) {
		return root
	}
	return nil
}

func pruneTreeDFS(root *TreeNode, parent *TreeNode, LR string) {
	if root == nil {
		return
	}
	pruneTreeDFS(root.Left, root, "L")
	pruneTreeDFS(root.Right, root, "R")
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		if LR == "L" {
			parent.Left = nil
		} else if LR == "R" {
			parent.Right = nil
		}
	}
}

func containsOne(root *TreeNode) bool {
	if root == nil {
		return false
	}
	Left := containsOne(root.Left)
	Right := containsOne(root.Right)
	if !Left {
		root.Left = nil
	}
	if !Right {
		root.Right = nil
	}
	return root.Val == 1 || Left || Right

}

func Test_pruneTree(t *testing.T) {
	root := &TreeNode{
		Val: 0,
	}
	data := pruneTree(root)
	print(data)
}
