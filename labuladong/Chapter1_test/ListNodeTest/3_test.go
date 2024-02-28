package ListNodeTest

import (
	"fmt"
	"strings"
	"testing"
)

//如何判断回文链表

func Test_3(t *testing.T) {
	fmt.Println(isPalindromeString("asddsa"))
	fmt.Println(isPalindromeString("asdbsa"))
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Println(isPalindromeListNode(head))
}

//判断一个字符串是不是回文串
func isPalindromeString(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		if s[l] < 'a' || s[l] > 'z' {
			l++
			continue
		}
		if s[r] < 'a' || s[r] > 'z' {
			l++
			continue
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

//一、判断回文单链表
//https://leetcode.cn/problems/palindrome-linked-list/description/
func isPalindromeListNode(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	var reverse func(head *ListNode) *ListNode
	reverse = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		last := reverse(head.Next)
		head.Next.Next = head
		head.Next = nil
		return last
	}
	node1 := head
	node2 := reverse(slow)
	for node2 != nil {
		if node1.Val != node2.Val {
			return false
		}
		node1 = node1.Next
		node2 = node2.Next
	}
	return true
}
