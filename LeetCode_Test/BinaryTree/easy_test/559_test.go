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
}

func maxDepth0(root *Node) (ans int) {
	if root == nil {
		return 0
	}
	ans = 1
	var helper func(root *Node, deep int)
	helper = func(root *Node, deep int) {
		if root == nil {
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
