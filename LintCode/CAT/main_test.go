package CAT

import (
	"fmt"
	"testing"
)

//
func Test(t *testing.T) {
	fmt.Println("hello test")
	//153,370,371,407
	//	getNarcissisticNumbers(1)
	//a := []int{1, 3, 1, 4, 4, 2}
	//fmt.Println(deduplication(a))
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	insertNode(head, 3)
}

//228. 链表的中点
func middleNode(head *ListNode) *ListNode {
	// write your code here
	dp := make(map[int]*ListNode)
	count := 0
	for head != nil {
		dp[count] = head
		count++
		head = head.Next
	}
	return dp[(count-1)/2]
}

//219. 在排序链表中插入一个节点
func insertNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return &ListNode{
			Val:  val,
			Next: nil,
		}
	}
	cur := head
	for cur != nil {
		//start
		if cur.Val > val {
			node := &ListNode{
				Val:  val,
				Next: head,
			}
			return node
		}
		//end
		if cur.Next == nil {
			cur.Next = &ListNode{
				Val:  val,
				Next: nil,
			}
			return head
		}
		//center
		var next *ListNode
		next = cur.Next
		if cur.Val < val && val < next.Val {
			temp := cur.Next
			cur.Next = &ListNode{
				Val:  val,
				Next: temp,
			}
			break
		}
		cur = cur.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//1663. 忧郁
func depress(m int, k int, arr []int) string {
	result := 0
	for i := 0; i < k; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		result += arr[i]
	}
	if result >= m {
		return "no"
	}
	return "yes"
}

func canAttendMeetings_1(intervals []*Interval) bool {
	//sort.Slice(intervals, func(i, j int) bool {
	//	return intervals[i].Start<intervals[j].Start
	//})
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i].Start > intervals[j].Start {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i].End > intervals[i+1].Start {
			return false
		}
	}
	return true
}
func canAttendMeetings(intervals []*Interval) bool {
	// Write your code here
	dp := make(map[int]int)
	for _, v := range intervals {
		interval := *v
		for i := interval.Start; i < interval.End; i++ {
			if dp[i] == 0 {
				dp[i] += 1
			} else {
				return false
			}
		}
	}
	return true
}

type Interval struct {
	Start, End int
}

//521. 去除重复元素
func deduplication(nums []int) int {
	dp := make(map[int]int)
	count := 0
	for i := 0; i < len(nums); i++ {
		if dp[nums[i]] == 0 {
			dp[nums[i]] = 1
		} else {
			count++
			temp := nums[i]
			nums[i] = nums[len(nums)-count]
			nums[len(nums)-count] = temp
			i--
		}
	}
	return len(nums) - count
}

//60. 搜索插入位置
func searchInsert_1(A []int, target int) int {
	if A == nil || len(A) == 0 || A[0] == 0 {
		return 0
	}
	left := 0
	right := len(A) - 1
	for left+1 < right {
		mid := (left + right) / 2
		if A[mid] > target {
			right = mid
		} else if A[mid] < target {
			left = mid
		} else {
			return mid
		}
	}
	if A[left] == target {
		return left
	} else if A[right] == target {
		return right
	} else if target > A[right] {
		return right + 1
	} else if target < A[left] {
		return left
	}
	return left + 1
}

func searchInsert(A []int, target int) int {
	if A == nil || len(A) == 0 || A[0] == 0 {
		return 0
	}
	for k, v := range A {
		if v >= target {
			return k
		}
	}
	return len(A)
}

//463. 整数排序
func sortIntegers(A *[]int) {
	// write your code here
	A_len := len(*A)
	for i := 0; i < A_len; i++ {
		for j := i + 1; j < A_len; j++ {
			if (*A)[i] > (*A)[j] {
				(*A)[i], (*A)[j] = (*A)[j], (*A)[i]
			}
		}
	}
}

//1654. 出现次数最多的字母
func mostFrequentlyAppearingLetters(str string) int {
	// Write your code here.
	dp := make(map[uint8]int, 52)
	result := 0
	for i := range str {
		dp[str[i]] += 1
		if dp[str[i]] > result {
			result = dp[str[i]]
		}
	}
	return result
}

func firstUniqChar(str string) byte {
	// Write your code here
	dp := make(map[int32]int, 0)
	for _, v := range str {
		dp[v] += 1
	}
	for _, v := range str {
		if dp[v] == 1 {
			return byte(v)
		}
	}
	return 0
}

func getNarcissisticNumbers(n int) []int {
	// write your code here
	result := make([]int, 0)
	//if n == 1 {
	//	for i := 0; i < 10; i++ {
	//		result = append(result, i)
	//	}
	//	return result
	//}
	min := pow(10, n-1)
	max := pow(10, n)
	for i := min; i < max; i++ {
		j := i
		s := 0
		if i == 153 {
			print(123)
		}
		for j > 0 {
			s += pow(j%10, n)
			j = j / 10
		}
		if s == i {
			result = append(result, i)
		}
	}
	return result
}

func pow(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}
