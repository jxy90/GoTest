package middle_test

import "testing"

//var index int
var index = 0

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseBetween(head, m, n)
	if index > m && index < n {
		head.Next.Next = head
		head.Next = nil
	}
	index++
	return p
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
	println(ans)
}
