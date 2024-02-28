package ListNodeTest

//双指针技巧秒杀七道链表题目

type ListNode struct {
	Val  int
	Next *ListNode
}

//单链表的分解
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	list1, list2 := &ListNode{}, &ListNode{}
	n1, n2, p := list1, list2, head
	for p != nil {
		if p.Val < x {
			n1.Next = p
			n1 = n1.Next
		} else {
			n2.Next = p
			n2 = n2.Next
		}
		p = p.Next
	}
	n2.Next = nil
	n1.Next = list2.Next
	return list1.Next
}

//合并 k 个有序链表
func mergeKLists(lists []*ListNode) *ListNode {
	var node *ListNode
	for i := range lists {
		node = merge2Lists(node, lists[i])
	}
	return node
}

func merge2Lists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	head := &ListNode{}
	tail := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	if list1 != nil {
		tail.Next = list1
	}
	if list2 != nil {
		tail.Next = list2
	}
	return head.Next
}

//单链表的倒数第 k 个节点
func findFromEnd(head *ListNode, k int) *ListNode {
	p1, p2 := head, head
	for i := 0; i < k; i++ {
		p2 = p2.Next
	}
	for p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}

//删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	node := findFromEnd(dummy, n+1)
	node.Next = node.Next.Next
	return dummy.Next
}

//单链表的中点
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

//判断链表是否包含环
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

//环形链表 II
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

//两个链表是否相交
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
