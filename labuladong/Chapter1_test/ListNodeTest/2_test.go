package ListNodeTest

import (
	"testing"
)

//递归魔法：反转单链表

func Test_2(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	reverseKGroup(head, 2)
}

//迭代反转链表
func reverse(a *ListNode) *ListNode {
	var pre, cur *ListNode
	pre, cur = nil, a
	if cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

// 反转区间 [a, b) 的元素，注意是左闭右开
func reverseAB(a, b *ListNode) *ListNode {
	var pre, cur *ListNode
	pre, cur = nil, a
	for cur != b {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

//https://leetcode.cn/problems/reverse-nodes-in-k-group/
//K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverseAB(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}
