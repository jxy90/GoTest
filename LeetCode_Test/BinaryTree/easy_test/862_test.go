package easy_test

import (
	"testing"
)

func Test_distanceK(t *testing.T) {
	root5 := &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 6},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 4},
		},
	}
	root1 := &TreeNode{
		Val:  3,
		Left: root5,
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 8},
		},
	}
	fmt.Println(distanceK(root1, root5, 2))
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	patents := make(map[int]*TreeNode)
	distanceKLoop(root, nil, patents)
	ans := make([]int, 0)
	distanceKLoop2(target, nil, patents, &ans, 0, k)
	return ans
}

func distanceKLoop(root, parent *TreeNode, parents map[int]*TreeNode) {
	if root == nil {
		return
	}
	parents[root.Val] = parent
	distanceKLoop(root.Left, root, parents)
	distanceKLoop(root.Right, root, parents)
}

func distanceKLoop2(root, target *TreeNode, parents map[int]*TreeNode, ans *[]int, deep, k int) {
	if root == nil {
		return
	}
	if deep == k {
		*ans = append(*ans, root.Val)
		return
	}
	if root.Left != target {
		distanceKLoop2(root.Left, root, parents, ans, deep+1, k)
	}
	if root.Right != target {
		distanceKLoop2(root.Right, root, parents, ans, deep+1, k)
	}
	if parents[root.Val] != target {
		distanceKLoop2(parents[root.Val], root, parents, ans, deep+1, k)
	}
}
