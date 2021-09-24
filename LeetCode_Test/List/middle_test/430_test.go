package middle_test

import "testing"

func Test_flatten(t *testing.T) {

}

func flatten(root *NodeWC) *NodeWC {
	//在呕吐节点前面再加一个节点,方便处理
	dummy := &NodeWC{}
	dummy.Next = root
	for root != nil {
		if root.Child != nil {
			//保存next节点
			temp := root.Next
			//得到扁平化的child
			child := flatten(root.Child)
			//root->child child->root
			root.Next = child
			child.Prev = root
			//清空child
			root.Child = nil
			//走完root
			for root.Next != nil {
				root = root.Next
			}
			//指向我们的next
			root.Next = temp
			if temp != nil {
				temp.Prev = root
			}
		}
		root = root.Next
	}
	return root.Next
}
