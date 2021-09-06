package easy_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	//return isSubtreeCheck(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
	tempS := isSubtreeLoop(s, "")
	tempT := isSubtreeLoop(t, "")
	fmt.Println(tempS)
	fmt.Println(tempT)
	return strings.Contains(tempS, tempT)
}

//暴力方式 O(s*t) O(s*t)
func isSubtreeCheck(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val == t.Val {
		return isSubtreeCheck(s.Left, t.Left) && isSubtreeCheck(s.Right, t.Right)
	}
	return false
}

//O(s+t) O(s+t)
func isSubtreeLoop(root *TreeNode, T string) string {
	if root == nil {
		return T + "nil"
	}
	return "," + strconv.Itoa(root.Val) + "," + isSubtreeLoop(root.Left, "L") + "," + isSubtreeLoop(root.Right, "R")
}

func Test_isSubtreeLoop(t *testing.T) {
	s := &TreeNode{
		Val: 12,
	}

	tt := &TreeNode{
		Val: 2,
	}
	isSubtree(s, tt)
}
