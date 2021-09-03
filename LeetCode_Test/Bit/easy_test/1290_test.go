package easy_test

import "testing"

func Test_getDecimalValue(t *testing.T) {
	fmt.Println(getDecimalValue(&ListNode{
		Val:  0,
		Next: nil,
	}))
}

func getDecimalValue(head *ListNode) int {
	ans := 0
	root := head
	for root != nil {
		ans = ans<<1 + root.Val
		root = root.Next
	}
	return ans
}

var ans = 0

func getDecimalValue0(head *ListNode) int {
	getDecimalValueHelper(head)
	return ans
}
func getDecimalValueHelper(head *ListNode) {
	if head == nil {
		return
	}
	ans = ans<<1 + head.Val
	getDecimalValueHelper(head.Next)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
