package easy_test

import (
	"sort"
	"strings"
	"testing"
)

func Test_0(t *testing.T) {
	//fmt.Println(findEvenNumbers([]int{2, 1, 3, 0}))
	//fmt.Println(findEvenNumbers([]int{2, 2, 8, 8, 2}))
	//fmt.Println(findEvenNumbers([]int{3, 7, 5}))
	//fmt.Println(findEvenNumbers([]int{0, 2, 0, 0}))
	//fmt.Println(findEvenNumbers([]int{0, 0, 0}))
	//fmt.Println(ListNode{Val: 1})
	//fmt.Println(ListNode{Val: 1})
	//node := TreeNode{
	//	Val:  2,
	//	Left: &TreeNode{Val: 1},
	//	Right: &TreeNode{
	//		Val:   3,
	//		Left:  &TreeNode{Val: 6},
	//		Right: &TreeNode{Val: 4},
	//	},
	//}
	//fmt.Println(getDirections(&node, 1, 6))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	str := make([]byte, 0)
	ss := make([]byte, 0)
	var find func(root *TreeNode, value int)
	find = func(root *TreeNode, value int) {
		if root == nil {
			return
		}
		if root.Val == value {
			for i := range str {
				ss = append(ss, str[i])
			}
			return
		}
		//root
		//left
		str = append(str, 'L')
		find(root.Left, value)
		str = str[:len(str)-1]
		//right
		str = append(str, 'R')
		find(root.Right, value)
		str = str[:len(str)-1]
	}
	ans := ""
	find(root, startValue)
	s1 := ss
	ss = make([]byte, 0)
	find(root, destValue)
	s2 := ss
	for len(s1) > 0 && len(s2) > 0 && s1[0] == s2[0] {
		s1 = s1[1:]
		s2 = s2[1:]
	}
	for i := len(s1); i > 0; i-- {
		ans += "U"
	}
	return strings.Repeat("U", len(s1)) + string(s2)
}

func validArrangement(pairs [][]int) [][]int {
	n := len(pairs)
	visited := map[int]bool{}
	flagFind := false
	temp := make([][]int, 0)
	ans := make([][]int, 0)
	cache := map[int]int{}
	var dfs func(begin, deep int)
	dfs = func(begin, deep int) {
		if len(temp) == n {
			for i := range temp {
				ans = append(ans, temp[i])
			}
			flagFind = true
			return
		}
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			if len(temp) == 0 || (len(temp) > 0 && temp[len(temp)-1][1] == pairs[i][0]) {
				if v, ok := cache[i]; ok && deep+v < n {
					cache[begin] = deep + v
					continue
				}
				visited[i] = true
				temp = append(temp, pairs[i])
				if begin == -1 {
					dfs(i, deep+1)
				} else {
					dfs(begin, deep+1)
				}
				if flagFind {
					return
				}
				visited[i] = false
				temp = temp[:len(temp)-1]
			}
		}
		cache[begin] = deep
	}
	dfs(-1, 1)
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	mem := head
	fast := head
	slow := head
	slowF := head
	for fast != nil && fast.Next != nil {
		slowF = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	slowF.Next = slowF.Next.Next
	return mem
}

func findEvenNumbers(digits []int) (ans []int) {
	n := len(digits)
	cache := map[int]bool{}
	visited := map[int]bool{}
	var dfs func(index int, temp []int)
	dfs = func(index int, temp []int) {
		if len(temp) == 3 {
			num := intSliceToNum(temp)
			if num%2 == 0 {
				cache[num] = true
			}
			return
		}
		for i := 0; i < n; i++ {
			if visited[i] == true {
				continue
			} else if len(temp) == 0 && digits[i] == 0 {
				continue
			} else if len(temp) == 2 && digits[i]%2 == 1 {
				continue
			}
			temp = append(temp, digits[i])
			visited[i] = true
			dfs(i+1, temp)
			temp = temp[:len(temp)-1]
			visited[i] = false
		}
	}
	dfs(0, make([]int, 0))
	//for i := 0; i < n; i++ {
	//	dfs(i, make([]int, 0))
	//}
	for k, _ := range cache {
		ans = append(ans, k)
	}
	sort.Ints(ans)
	return ans
}

func intSliceToNum(slice []int) (num int) {
	for _, i := range slice {
		num = num*10 + i
	}
	return num
}
