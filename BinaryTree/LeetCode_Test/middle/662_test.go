package middle_test

import "testing"

type ANode struct {
	Node     *TreeNode
	Position int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queueANode := []*ANode{&ANode{
		Node:     root,
		Position: 1,
	}}
	level := 0
	maxWidth := 0
	for len(queueANode) > 0 {
		length := len(queueANode)
		for i := 0; i < length; i++ {
			aNode := queueANode[i]
			width := queueANode[i].Position - queueANode[0].Position + 1
			if width > maxWidth {
				maxWidth = width
			}
			if aNode.Node.Left != nil {
				queueANode = append(queueANode, &ANode{
					Node:     aNode.Node.Left,
					Position: 2 * aNode.Position,
				})
			}
			if aNode.Node.Right != nil {
				queueANode = append(queueANode, &ANode{
					Node:     aNode.Node.Right,
					Position: 2*aNode.Position + 1,
				})
			}
		}
		level++
		queueANode = queueANode[length:]
	}
	return maxWidth
}

func Test_widthOfBinaryTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	widthOfBinaryTree(root)
}
