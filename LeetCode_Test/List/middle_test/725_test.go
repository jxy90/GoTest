package middle_test

import (
	"fmt"
	"testing"
)

func Test_splitListToParts(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	fmt.Println(splitListToParts(head, 5))
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	//取长度
	root := head
	length := 0
	for head != nil {
		head = head.Next
		length++
	}
	//每一份的长度
	//余数
	per := length / k
	yu := length % k
	//每一份
	ans := make([]*ListNode, k)
	head = root
	cnt := 0
	for i := 0; i < k; i++ {
		temp := head
		if yu > 0 {
			yu--
			temp = head
			head = head.Next
		}
		for head != nil && per != 0 {
			temp = head
			head = head.Next
			cnt++
			//长度满足|数据取完
			if cnt == per || head == nil {
				cnt = 0
				break
			}
		}
		if temp != nil {
			temp.Next = nil
		}
		ans[i] = root
		root = head
	}
	return ans
}
