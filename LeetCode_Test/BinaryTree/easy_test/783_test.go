package easy_test

import (
	"fmt"
	"math"
	"testing"
)

//func minDiffInBST(root *TreeNode) int {
//	nums := minDiffInBSTHelper(root)
//	min := math.MaxInt32
//	for i := range nums {
//		if i > 0 {
//			min = CommonUtil.Min(min, nums[i]-nums[i-1])
//		}
//	}
//	return min
//}
//
//func minDiffInBSTHelper(root *TreeNode) []int {
//	var temp []int
//	if root == nil {
//		return temp
//	}
//	//left
//	if root.Left != nil {
//		temp = append(temp, minDiffInBSTHelper(root.Left)...)
//	}
//	//c
//	temp = append(temp, root.Val)
//	//right
//	if root.Right != nil {
//		temp = append(temp, minDiffInBSTHelper(root.Right)...)
//	}
//	return temp
//}

func minDiffInBST(root *TreeNode) int {
	minDiffInBST2(root)
	return min
}

var min = math.MaxInt32
var pre = -1

func minDiffInBST2(root *TreeNode) {
	if root == nil {
		return
	}
	minDiffInBST2(root.Left)
	if pre != -1 && root.Val-pre < min {
		min = root.Val - pre
	}
	pre = root.Val
	minDiffInBST2(root.Right)
}

func Test_minDiffInBST(t *testing.T) {
	//root := &TreeNode{
	//	Val: 4,
	//	Left: &TreeNode{
	//		Val:   2,
	//		Left:  &TreeNode{Val: 1},
	//		Right: &TreeNode{Val: 3},
	//	},
	//	Right: &TreeNode{
	//		Val: 6,
	//	},
	//}

	//[27,null,34,null,58,50,null,44]
	//root := &TreeNode{
	//	Val: 27,
	//	Right: &TreeNode{
	//		Val: 34,
	//		Right: &TreeNode{
	//			Val: 58,
	//			Left: &TreeNode{
	//				Val:  50,
	//				Left: &TreeNode{Val: 44},
	//			},
	//		},
	//	},
	//}

	//[1,0,48,null,null,12,49]
	root1 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 0},
		Right: &TreeNode{
			Val:   48,
			Left:  &TreeNode{Val: 12},
			Right: &TreeNode{Val: 49},
		},
	}
	data := minDiffInBST(root1)
	fmt.Println(data)
}
