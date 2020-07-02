package easy_test

import (
	"strconv"
	"testing"
)

var str1, str2 string

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	str1, str2 = "", ""
	leafSimilarDFS(root1, "1")
	leafSimilarDFS(root2, "2")
	return str1 == str2
}

func leafSimilarDFS(root *TreeNode, str string) {
	if root == nil {
		return
	}
	leafSimilarDFS(root.Left, str)
	if root.Right == nil && root.Left == nil {
		if str == "1" {
			str1 += "," + strconv.Itoa(root.Val)
		} else if str == "2" {
			str2 += "," + strconv.Itoa(root.Val)
		}
	}
	leafSimilarDFS(root.Right, str)
}

func Test_leafSimilar(t *testing.T) {
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	root2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
		},
	}
	data := leafSimilar(root1, root2)
	println(data)
}
