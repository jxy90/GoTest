package middle_test

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	hasCycle := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			hasCycle = true
			break
		}
	}
	if !hasCycle {
		return nil
	}
	slow = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
