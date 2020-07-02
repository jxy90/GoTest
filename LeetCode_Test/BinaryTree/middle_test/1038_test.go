package middle_test

import "testing"

var sum int

func bstToGst(root *TreeNode) *TreeNode {
	sum = 0
	bstToGstDFS(root)
	return root
}

func bstToGstDFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	bstToGstDFS(root.Right)
	sum += root.Val
	root.Val = sum
	bstToGstDFS(root.Left)
	return root
}

func Test_bstToGst(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 7,
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}
	data := bstToGst(root)
	println(data)
}
