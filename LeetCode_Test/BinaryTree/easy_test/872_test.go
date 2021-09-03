package easy_test

import (
	"strconv"
	"testing"
)

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
	fmt.Println(data)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	nums1 := make([]int, 0)
	nums2 := make([]int, 0)
	leafSimilarInOrder(root1, &nums1)
	leafSimilarInOrder(root2, &nums2)
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

func leafSimilarInOrder(node *TreeNode, nums *[]int) {
	if node == nil {
		return
	}
	leafSimilarInOrder(node.Left, nums)
	if node.Left == nil && node.Right == nil {
		*nums = append(*nums, node.Val)
		return
	}
	leafSimilarInOrder(node.Right, nums)
}

var str1, str2 string

func leafSimilar0(root1 *TreeNode, root2 *TreeNode) bool {
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
