package Tree

import (
	"fmt"
	"testing"
)

//最大二叉树
//https://leetcode.cn/problems/maximum-binary-tree/description/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := 0
	for i, v := range nums {
		if nums[index] < v {
			index = i
		}
	}
	left := constructMaximumBinaryTree(nums[:index])
	right := constructMaximumBinaryTree(nums[index+1:])
	return &TreeNode{
		Val:   nums[index],
		Left:  left,
		Right: right,
	}
}

//从前序与中序遍历序列构造二叉树
//https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
//func buildTree(preorder []int, inorder []int) *TreeNode {
//	return buildTreeHelper(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
//}
//
//func buildTreeHelper(preorder []int, inorder []int, preLeft, preRight, inLeft, inRight int) *TreeNode {
//	if preLeft > preRight {
//		return nil
//	}
//
//	rootVal := preorder[preLeft]
//	rootIndex := 0
//	for i := inLeft; i <= inRight; i++ {
//		if inorder[i] == rootVal {
//			rootIndex = i
//			break
//		}
//	}
//	leftSize := rootIndex - inLeft
//	return &TreeNode{
//		Val:   rootVal,
//		Left:  buildTreeHelper(preorder, inorder, preLeft+1, preLeft+leftSize, inLeft, rootIndex),
//		Right: buildTreeHelper(preorder, inorder, preLeft+leftSize+1, preRight, rootIndex+1, inRight),
//	}
//
//}

//从中序与后序遍历序列构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	return buildTreeHelper(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

func buildTreeHelper(inorder []int, postorder []int, inLeft, inRight, postLeft, postRight int) *TreeNode {
	if postLeft > postRight {
		return nil
	}

	rootVal := postorder[postRight]
	rootIndex := 0
	for i := inLeft; i <= inRight; i++ {
		if rootVal == inorder[i] {
			rootIndex = i
			break
		}
	}
	leftSize := rootIndex - inLeft
	return &TreeNode{
		Val:   rootVal,
		Left:  buildTreeHelper(inorder, postorder, inLeft, rootIndex-1, postLeft, postLeft+leftSize-1),
		Right: buildTreeHelper(inorder, postorder, rootIndex+1, inRight, postLeft+leftSize, postRight-1),
	}
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	return constructFromPrePostHelper(preorder, 0, len(preorder)-1, postorder, 0, len(postorder)-1)
}

func constructFromPrePostHelper(preorder []int, preStart, preEnd int, postorder []int, postStart, postEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	if preStart == preEnd {
		return &TreeNode{Val: preorder[preEnd]}
	}
	rootVal := preorder[preStart]
	leftVal := preorder[preStart+1]
	leftIndex := 0
	for i := postStart; i < postEnd+1; i++ {
		if leftVal == postorder[i] {
			leftIndex = i
		}
	}
	leftSize := leftIndex - postStart + 1
	return &TreeNode{
		Val:   rootVal,
		Left:  constructFromPrePostHelper(preorder, preStart+1, preStart+leftSize, postorder, postStart, leftIndex),
		Right: constructFromPrePostHelper(preorder, preStart+leftSize+1, preEnd, postorder, leftIndex+1, postEnd-1),
	}
}

func Test_2(t *testing.T) {
	fmt.Println(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}
