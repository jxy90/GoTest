package main_test

import (
	"fmt"
	"testing"
)

func (bt *TreeNode) preorderTraversal() []int {
	//中-》左-》右
	var temp []int
	if bt == nil {
		return temp
	}
	temp = append(temp, bt.Val)
	if bt.Left != nil {
		temp = append(temp, bt.Left.preorderTraversal()...)
	}
	if bt.Right != nil {
		temp = append(temp, bt.Right.preorderTraversal()...)
	}
	return temp
}

//非递归
func (root *TreeNode) preorderTraversalNew() []int {
	if root == nil {
		return nil
	}
	var stack []*TreeNode
	var result []int
	temp := root
	//中-》左-》右
	for temp != nil || len(stack) != 0 {
		if temp != nil {
			result = append(result, temp.Val)
			stack = append(stack, temp)
			temp = temp.Left
		} else {
			temp = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}
	return result
}

func Test_preorderTraversalMirror(t *testing.T) {

	root10 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(root10.preorderTraversalMirror())
}

//https://www.cnblogs.com/blzm742624643/p/10021388.html
//https://www.cnblogs.com/floridezzh/p/4899347.html
//mirro算法
func (root *TreeNode) preorderTraversalMirror() []int {
	if root == nil {
		return nil
	}
	cur := root
	var mostRight *TreeNode
	result := make([]int, 0)
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur
				result = append(result, cur.Val)
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
			}
		} else {
			result = append(result, cur.Val)
		}
		cur = cur.Right
	}
	return result
}
