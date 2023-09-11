package middle_test

import (
	"fmt"
	"testing"
)

func Test_2(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	fmt.Printf("%v", addTwoNumbers(l1, l2))
	l1 = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 9,
			},
		},
	}
	l2 = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}
	//10407
	fmt.Printf("%v", addTwoNumbers(l1, l2))
	l1 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 1,
							},
						},
					},
				},
			},
		},
	}
	l2 = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	//10407
	fmt.Printf("%v", addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		sum = sum % 10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	n1, n2 := list2num(l1), list2num(l2)
	fmt.Println(n1, n2)
	fmt.Println(n1 + n2)
	return num2list(n1 + n2)
}

func list2num(l *ListNode) int64 {
	if l.Next == nil {
		return int64(l.Val)
	}
	num := int64(0)
	if l.Next != nil {
		num = list2num(l.Next)*10 + int64(l.Val)
	}
	return num
}

func num2list(n int64) *ListNode {
	if n < 10 {
		return &ListNode{Val: int(n)}
	}

	temp := n % 10
	n = n / 10
	return &ListNode{
		Val:  int(temp),
		Next: num2list(n),
	}
}
