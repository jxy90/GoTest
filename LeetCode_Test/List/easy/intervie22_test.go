package easy_test

import "testing"

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func Test_getKthFromEnd(t *testing.T) {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	getKthFromEnd(head, 1)
}
