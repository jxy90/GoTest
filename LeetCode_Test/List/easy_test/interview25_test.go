package easy_test

import "testing"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}

func Test_mergeTwoLists(t *testing.T) {
	l1 := &ListNode{
		Val: 0,
	}
	l2 := &ListNode{
		Val: 0,
	}
	mergeTwoLists(l1, l2)
}
