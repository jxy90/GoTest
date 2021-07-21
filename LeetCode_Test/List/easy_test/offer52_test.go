package easy_test

func getIntersectionNode00(headA, headB *ListNode) *ListNode {
	cacheA := map[*ListNode]bool{}
	for headA != nil || headB != nil {

		if headA != nil {
			if _, ok := cacheA[headA]; ok {
				return headA
			}
			cacheA[headA] = true
			headA = headA.Next
		}
		if headB != nil {
			if _, ok := cacheA[headB]; ok {
				return headB
			}
			cacheA[headB] = true
			headB = headB.Next
		}
	}
	return nil
}
