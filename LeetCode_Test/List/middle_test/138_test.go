package middle_test

import (
	"fmt"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node1.Next = node2
	node1.Random = node2
	node2.Random = node2
	fmt.Println(copyRandomList(node1))
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	//保存head和headCopy的map
	copied := make(map[*Node]*Node, 0)
	node := head
	//回溯得到复制(除去random)
	headCopy := loop(copied, node)

	nodeCopy := headCopy
	node = head
	//循环添加random
	for node != nil {
		if node.Random != nil {
			if _, ok := copied[node.Random]; ok {
				nodeCopy.Random = copied[node.Random]
			}
		}
		nodeCopy = nodeCopy.Next
		node = node.Next
	}

	return headCopy
}

func loop(copied map[*Node]*Node, head *Node) *Node {
	if head == nil {
		return nil
	}
	next := loop(copied, head.Next)
	headCopy := &Node{
		Val:  head.Val,
		Next: next,
		//Random: nil,
	}
	copied[head] = headCopy
	return headCopy
}
