package middle_test

import "testing"

type CBTInserter struct {
	Root  *TreeNode
	Nodes []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	result := make([]*TreeNode, 0)
	queue := make([]*TreeNode, 0)
	result = append(result, root)
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
				result = append(result, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
				result = append(result, queue[i].Right)
			}
		}
		queue = queue[length:]
	}
	return CBTInserter{
		Root:  root,
		Nodes: result,
	}
}

func (this *CBTInserter) Insert(v int) int {
	node := &TreeNode{
		Val: v,
	}
	this.Nodes = append(this.Nodes, node)
	length := len(this.Nodes)
	parent := this.Nodes[length/2-1]
	if length%2 == 1 {
		parent.Right = node
	} else {
		parent.Left = node
	}
	return parent.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.Root
}

func Test_CBTInserter(*testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: nil,
		},
	}
	CBT := Constructor(root)
	CBT.Insert(7)
	CBT.Insert(8)
	CBT.Get_root()
}
