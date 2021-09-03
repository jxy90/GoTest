package easy_test

import "testing"

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

func reverseListNew(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseListNew(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

var temp *ListNode

func reverseListN(head *ListNode, N int) *ListNode {
	if N == 1 {
		temp = head.Next
		return head
	}
	p := reverseListN(head.Next, N-1)
	head.Next.Next = head
	head.Next = temp
	return p
}

func Test_reverseList(t *testing.T) {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					//Next: &ListNode{
					//	Val: 5,
					//	Next: &ListNode{
					//		Val:  6,
					//		Next: nil,
					//	},
					//},
				},
			},
		},
	}
	last := reverseListN(node, 3)
	print(last)
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == 1 {
		return reverseListN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

func Test_reverseBetween(t *testing.T) {
	nums := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	n, m := 2, 4
	ans := reverseBetween(nums, n, m)
	fmt.Println(ans)
}
