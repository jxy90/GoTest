package easy_test

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

func getIntersectionNode0(headA, headB *ListNode) *ListNode {
	f := make(map[*ListNode]bool)
	for headA != nil || headB != nil {
		if headA != nil {
			if _, ok := f[headA]; ok {
				return headA
			}
			f[headA] = true
			headA = headA.Next
		}
		if headB != nil {
			if _, ok := f[headB]; ok {
				return headB
			}
			f[headB] = true
			headB = headB.Next
		}
	}
	return nil
}
