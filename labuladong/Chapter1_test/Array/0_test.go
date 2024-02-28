package Array

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//双指针技巧秒杀七道数组题目

//一、快慢指针技巧
//https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

//删除排序链表中的重复元素」
//https://leetcode.cn/problems/remove-duplicates-from-sorted-list/description/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil {
		if slow.Val != fast.Val {
			slow = slow.Next
			slow.Val = fast.Val
		}
		fast = fast.Next
	}
	slow.Next = nil
	return head
}

//移除元素
//https://leetcode.cn/problems/remove-element/description/
func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
	return slow
}

//移动零」
//https://leetcode.cn/problems/move-zeroes/description/
func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}

//二、左右指针的常用算法
//1、二分查找
//子目录二分查找

//2、两数之和
//子目录nSum

//3、反转数组
//https://leetcode.cn/problems/reverse-string/
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

//4、回文串判断
//https://leetcode.cn/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	ans := ""
	for i := range s {
		//i 为中心的最长回文串
		s1 := palindrome(s, i, i)
		//i 和 i +1 为中心的最长回文串
		s2 := palindrome(s, i, i+1)
		if len(s1) > len(ans) {
			ans = s1
		}
		if len(s2) > len(ans) {
			ans = s2
		}
	}

	return ans
}

func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

func Test_0(t *testing.T) {
	//fmt.Println(isPalindromeN("bb", 2))
	fmt.Println(longestPalindrome("aba"))
}
