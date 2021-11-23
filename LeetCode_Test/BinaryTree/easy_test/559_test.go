package easy_test

import (
	"fmt"
	"testing"
)

func Test_maxDepth0(t *testing.T) {
	s := &Node{
		Val: 12,
	}
	fmt.Println(maxDepth0(s))
	s = &Node{
		Val:      12,
		Children: nil,
	}
	fmt.Println(maxDepth0(s))
}

func maxDepth1(root *Node) (ans int) {
	if root == nil {
		return 0
	}
	ans = 1
	var helper func(root *Node, deep int)
	helper = func(root *Node, deep int) {
		if root.Children == nil {
			return
		}
		for _, child := range root.Children {
			helper(child, deep+1)
			if deep+1 > ans {
				ans = deep + 1
			}
		}
	}
	helper(root, 1)
	return ans
}

func maxDepth0(root *Node) (ans int) {
	var helper func(root *Node, deep int) int
	helper = func(root *Node, deep int) int {
		if root == nil {
			return 0
		}
		if root.Children == nil {
			return 1
		}
		max := 1
		for _, child := range root.Children {
			temp := helper(child, deep) + 1
			if temp > max {
				max = temp
			}
		}
		return max
	}
	return helper(root, 1)
}
