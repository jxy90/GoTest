package middle_test

import (
	"fmt"
	"testing"
)

func Test_isEvenOddTree(t *testing.T) {
	//root := &TreeNode{
	//	Val: 5,
	//	Left: &TreeNode{
	//		Val: 4,
	//		Left: &TreeNode{
	//			Val: 3,
	//		},
	//		Right: &TreeNode{
	//			Val: 3,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val: 7,
	//		},
	//		Right: nil,
	//	},
	//}

	//root := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 10,
	//		Left: &TreeNode{
	//			Val:   3,
	//			Left:  &TreeNode{Val: 12},
	//			Right: &TreeNode{Val: 8},
	//		},
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val: 4,
	//		Left: &TreeNode{
	//			Val:   7,
	//			Left:  &TreeNode{Val: 6},
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   9,
	//			Left:  nil,
	//			Right: &TreeNode{Val: 2},
	//		},
	//	},
	//}

	//root := &TreeNode{
	//	Val: 5,
	//	Left: &TreeNode{
	//		Val: 9,
	//		Left: &TreeNode{
	//			Val: 3,
	//		},
	//		Right: &TreeNode{
	//			Val: 5,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val: 1,
	//		Left: &TreeNode{
	//			Val: 7,
	//		},
	//		Right: nil,
	//	},
	//}
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 12,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 18,
				},
				Right: &TreeNode{
					Val: 16,
				},
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
		Right: &TreeNode{
			Val:   8,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(isEvenOddTree(root))
}

func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	cache := map[int][]*TreeNode{}
	level0 := make([]*TreeNode, 1)
	level0[0] = root
	cache[0] = level0
	level := 0
	for len(cache) > 0 {
		for len(cache[level]) > 0 {
			cur := cache[level][0]
			cache[level] = cache[level][1:]
			if cur.Left != nil && cur.Right != nil && cache[level+1] == nil {
				cache[level+1] = make([]*TreeNode, 0)
			}
			if cur.Left != nil {
				cache[level+1] = append(cache[level+1], cur.Left)
				if !isEvenOddTreeCheck(level+1, cache) {
					return false
				}
			}
			if cur.Right != nil {
				cache[level+1] = append(cache[level+1], cur.Right)
				if !isEvenOddTreeCheck(level+1, cache) {
					return false
				}
			}
		}
		delete(cache, level)
		level++
	}
	return true
	//return helper(root,0)
}

func isEvenOddTreeCheck(level int, cache map[int][]*TreeNode) bool {
	slice := cache[level]
	n := len(slice)
	if n > 0 && (level%2)^(slice[n-1].Val%2) == 0 {
		return false
	}
	if n < 2 {
		return true
	}
	if level%2 == 0 && slice[n-1].Val > slice[n-2].Val {
		return true
	}
	if level%2 == 1 && slice[n-1].Val < slice[n-2].Val {
		return true
	}
	return false
}

func isEvenOddTreeHelper(root *TreeNode, level int) bool {
	if root == nil {
		return true
	}
	return (level%2)^(root.Val%2) == 1 && isEvenOddTreeHelper(root.Left, level+1) && isEvenOddTreeHelper(root.Right, level+1)
}
