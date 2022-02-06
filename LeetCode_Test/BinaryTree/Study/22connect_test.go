package main_test

import "fmt"

//https://leetcode-cn.com/explore/learn/card/data-structure-binary-tree/4/conclusion/17/
//https://www.jianshu.com/p/2a7747f56b3a	-参考解法2
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		root.Left.Next = root.Right
	}
	if root.Right != nil && root.Next != nil {
		root.Right.Next = root.Next.Left
	}
	connect(root.Left)
	connect(root.Right)
	return root
}

func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Val == 9 {
		fmt.Println("---")
	}
	if root.Val == 7 {
		fmt.Println("---")
	}
	if root.Left != nil {
		if root.Right != nil {
			root.Left.Next = root.Right
		} else {
			temp := root.Next
			for temp != nil {
				if temp.Left != nil {
					root.Left.Next = temp.Left
					break
				}
				if temp.Right != nil {
					root.Left.Next = temp.Right
					break
				}
				temp = temp.Next
			}
		}
	}
	if root.Right != nil {
		temp := root.Next
		for temp != nil {
			if temp.Left != nil {
				root.Right.Next = temp.Left
				break
			}
			if temp.Right != nil {
				root.Right.Next = temp.Right
				break
			}
			temp = temp.Next
		}
	}
	connect2(root.Right)
	connect2(root.Left)
	return root
}
