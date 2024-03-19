package _022

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_baidu1(t *testing.T) {
	//root := &TreeNode{
	//	Val:   0,
	//	Right: nil,
	//	Left:  nil,
	//}
	////topToDown(root)
	//root = &TreeNode{
	//	Val: 0,
	//	Right: &TreeNode{
	//		Val: 2,
	//		Right: &TreeNode{
	//			Val:   6,
	//			Right: nil,
	//			Left:  nil,
	//		},
	//		Left: &TreeNode{
	//			Val:   5,
	//			Right: nil,
	//			Left:  nil,
	//		},
	//	},
	//	Left: &TreeNode{
	//		Val: 1,
	//		Right: &TreeNode{
	//			Val:   4,
	//			Right: nil,
	//			Left:  nil,
	//		},
	//		Left: &TreeNode{
	//			Val:   3,
	//			Right: nil,
	//			Left:  nil,
	//		},
	//	},
	//}
	//topToDown(root)

	fmt.Println(maxLength("abcde", "bcd"))
}

type TreeNode struct {
	Val         int
	Right, Left *TreeNode
}

//二叉树层次遍历
func topToDown(root *TreeNode) {
	if root == nil {
		return
	}
	//1.
	//q := make([]map[int]*TreeNode, 0)
	levelQueue := map[int][]*TreeNode{}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	levelQueue[0] = q
	level := 0
	for len(levelQueue) > 0 {
		for len(levelQueue[level]) > 0 {
			cur := levelQueue[level][0]
			levelQueue[level] = levelQueue[level][1:]
			fmt.Println(cur.Val)
			if (cur.Left != nil || cur.Right != nil) && levelQueue[level+1] == nil {
				levelQueue[level+1] = make([]*TreeNode, 0)
			}
			if cur.Left != nil {
				levelQueue[level+1] = append(levelQueue[level+1], cur.Left)
			}
			if cur.Right != nil {
				levelQueue[level+1] = append(levelQueue[level+1], cur.Right)
			}
		}
		delete(levelQueue, level)
		level++
	}
}

//mirros

//s1,s2最长公共子串
//abcdef
//babcd
func longestCommonSubsequence(text1 string, text2 string) int {
	//f[i][j]表示text1的i位置和text2的j位置,最长公共子串的长度
	m, n := len(text1), len(text2)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text1[j-1] {
				f[i][j] = f[i-1][j-1] + 1
				ans = CommonUtil.Max(ans, f[i][j])
			}
		}
	}
	for i := range f {
		fmt.Println(f[i])
	}
	return f[m][n]
}

func maxLength(s1, s2 string) int {
	m, n := len(s1), len(s2)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				f[i][j] = f[i-1][j-1] + 1
				ans = CommonUtil.Max(ans, f[i][j])
			} else {
				f[i][j] = 0
			}
		}
	}
	for i := range f {
		fmt.Println(f[i])
	}
	return ans
}

type Node struct {
	val  int
	next *Node
}

func getFirstNodeHaveRing(root *Node) *Node {
	//如果存在环，打印第一个节点
	if root == nil {
		return nil
	}
	cache := map[*Node]struct{}{}
	cache[root] = struct{}{}
	for root.next != nil {
		next := root.next
		if _, ok := cache[next]; ok {
			return next
		}
		cache[next] = struct{}{}
		root = next
	}
	return nil
}
