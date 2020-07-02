package middle_test

import "testing"

func sortList(head *ListNode) *ListNode {
	//边界情况
	if head == nil || head.Next == nil {
		return head
	}
	//快慢指针，找到中点实现二分
	slow, fast := head, head
	leftend := head
	for fast != nil && fast.Next != nil {
		leftend = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	right := slow
	left := head
	leftend.Next = nil
	left = sortList(left)
	right = sortList(right)

	return sortList_Merge(left, right)
}

//归并两个链表
func sortList_Merge(left *ListNode, right *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	cur := head
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return head.Next
}

func Test_sortList(t *testing.T) {
	node := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}
	data := sortList(node)
	println(data)
}
