package middle_test

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

type NodeWC struct {
	Val   int
	Prev  *NodeWC
	Next  *NodeWC
	Child *NodeWC
}
