package _022

import (
	"fmt"
	"testing"
)

func Test_mx(t *testing.T) {
	//fmt.Println(FindFirstK([]int{0, 1, 2, 2, 3, 3, 5}, 3)) //4
	//fmt.Println(FindFirstK([]int{0, 0, 0}, 0))             //0
	fmt.Println(FindFirstK([]int{0, 0, 0}, 1)) //0
	root := &Node1{
		Val: 1,
		Left: &Node1{
			Val: 2,
			Left: &Node1{
				Val: 4,
			},
			Right: nil,
		},
		Right: &Node1{
			Val:  3,
			Left: nil,
			Right: &Node1{
				Val: 5,
			},
		},
	}
	fmt.Println(TopToDown1(root))
}

//先序遍历: G DA FE M HZ
//中序遍历: A D EF G H M Z
//后续遍历：AEFDHZMG
//		G
//	D		M
//A  F	  H	  Z
//	E

//二叉树的层序遍历
type Node1 struct {
	Val         int
	Left, Right *Node1
}

func TopToDown1(root *Node1) (ans []int) {
	if root == nil {
		return
	}
	queue := make([]*Node1, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		ans = append(ans, cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return
}

//有序数组，有重复元素，找到一个给定值的第一个下标,找不到的时候返回-1
func FindFirstK(nums []int, k int) int {
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		mid := l + (r-l)>>1
		midVal := nums[mid]
		if midVal < k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if nums[l] != k {
		return -1
	}
	return l
}
