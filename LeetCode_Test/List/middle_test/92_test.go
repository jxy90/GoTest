package middle_test

import (
	"fmt"
	"testing"
)

//var index int
var index = 0

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	index++
	p := reverseBetween(head.Next, m, n)
	if index >= m && index < n {
		head.Next.Next = head
		head.Next = nil
	}
	index--
	return p
}

func reverseBetween2(head *ListNode, m int, n int) *ListNode {
	var ans *ListNode
	cur := head
	if cur != nil {
		temp := cur.Next
		cur.Next = ans
		ans = cur
		cur = temp
	}
	return ans
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
