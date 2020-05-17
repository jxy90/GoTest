package main

import (
	Codec2 "github.com/jxy90/GoTest/BinaryTree/Study/Codec"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Constructor
func NewBinaryTree(data int) *TreeNode {
	return &TreeNode{Val: data}
}

func main() {
	root := NewBinaryTree(1)
	root.Left = NewBinaryTree(2)
	root.Right = NewBinaryTree(2)
	root.Right.Left = NewBinaryTree(3)

	for _, v := range postorderTraversal(root) {
		print(v)
	}
	println("")

	arr := levelOrder(root)
	for _, v := range arr {
		for _, vv := range v {
			print(vv)
		}
	}
	println("level1")
	arr2 := levelLpoop(root, 0)
	for _, v := range arr2 {
		for _, vv := range v {
			print(vv)
		}
	}
	println("level2")

	println(TopDown(root, 0))
	println(BottomUp(root))
	//对称二叉树
	println(isSymmetric(root))
	//路径和
	println(hasPathSum(root, 6))

	//中序遍历 inorder = [9,3,15,20,7]
	//后序遍历 postorder = [9,15,7,20,3]
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	bt1 := buildTree(inorder, postorder)
	println(bt1)

	//前序遍历 preorder = [3,9,20,15,7]
	//中序遍历 inorder = [9,3,15,20,7]
	preorder := []int{3, 9, 20, 15, 7}
	inorder2 := []int{9, 3, 15, 20, 7}
	bt2 := buildTree2(preorder, inorder2)
	println(bt2)

	//填充每个节点的下一个右侧节点指针 II
	bt22 := &Node{Val: 2}
	bt22.Left = &Node{Val: 1}
	bt22.Right = &Node{Val: 3}
	bt22.Left.Left = &Node{Val: 0}
	bt22.Left.Right = &Node{Val: 7}
	bt22.Right.Left = &Node{Val: 9}
	bt22.Right.Right = &Node{Val: 1}
	bt22.Left.Left.Left = &Node{Val: 2}
	bt22.Left.Right.Left = &Node{Val: 1}
	bt22.Left.Right.Right = &Node{Val: 0}
	bt22.Left.Right.Right.Left = &Node{Val: 7}
	bt22.Right.Right.Left = &Node{Val: 8}
	bt22.Right.Right.Right = &Node{Val: 8}
	connect2(bt22)
	println(bt2)

	//obj := Constructor()
	//data := obj.serialize(root)
	////data2 := obj.serializeLoop2(root)
	//ans := obj.deserialize(data)
	//println(data, ans)
	//////////////////

	//[1,2,3,nil,nil,4,5]
	root10 := &Codec2.TreeNode{
		Val: 1,
		Left: &Codec2.TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &Codec2.TreeNode{
			Val: 3,
			Left: &Codec2.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &Codec2.TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}
	serializeStr := Codec2.Serialize(root10)
	println(serializeStr)
	deserializeObj := Codec2.Deserialize(serializeStr)
	println(deserializeObj)
}

var result2 [][]string

func levelLoop(node *TreeNode, level int) {
	if node == nil {
		return
	}
	if result2 == nil {
		result2 = make([][]string, 0)
	}
	if len(result2) <= level {
		result2 = append(result2, make([]string, 0))
	}
	result2[level] = append(result2[level], strconv.Itoa(node.Val))
	levelLoop(node.Left, level+1)
	levelLoop(node.Right, level+1)
}
