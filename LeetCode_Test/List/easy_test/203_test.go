package easy_test

func removeElements(head *ListNode, val int) *ListNode {
	head = &ListNode{
		Val:  0,
		Next: head,
	}
	node := head.Next
	pre := head
	for node != nil {
		if node.Val == val {
			node = node.Next
			pre.Next = node
		} else {
			node = node.Next
			pre = pre.Next
		}
	}
	return head.Next
}
