package easy_test

import (
	"testing"
)

var point *TreeNode
var temp *TreeNode

func convertBiNode(root *TreeNode) *TreeNode {
	point = &TreeNode{0, nil, nil}
	temp = point
	convertBiNodeDFS(root)
	return point.Right
}
func convertBiNodeDFS(root *TreeNode) {
	if root != nil {
		convertBiNodeDFS(root.Left)
		convertBiNodeVisit(root)
		convertBiNodeDFS(root.Right)
	}
}

func convertBiNodeVisit(root *TreeNode) {
	temp.Right = &TreeNode{
		Val: root.Val,
	}
	temp = temp.Right
}

func convertBiNode2(root *TreeNode) *TreeNode {
	var ans *TreeNode
	temp := ans
	preOrder := convertBiNodeHelper(root)
	for index := range preOrder {
		if index == 0 {
			ans = &TreeNode{
				Val: preOrder[index],
			}
			temp = ans
		} else {
			temp.Right = &TreeNode{
				Val: preOrder[index],
			}
			temp = temp.Right
		}
	}
	return ans
}

func convertBiNodeHelper(root *TreeNode) []int {
	var temp []int
	if root == nil {
		return temp
	}
	if root.Left != nil {
		temp = append(temp, convertBiNodeHelper(root.Left)...)
	}
	temp = append(temp, root.Val)
	if root.Right != nil {
		temp = append(temp, convertBiNodeHelper(root.Right)...)
	}
	return temp
}

func Test_convertBiNode(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	ans := convertBiNode(root)
	fmt.Println(ans)
}
