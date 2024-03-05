package Tree

import (
	"fmt"
	"testing"
)

func nextGreaterElement0(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	stack := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return res
}

//下一个更大元素 I
//https://leetcode.cn/problems/next-greater-element-i/description/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	temps := nextGreaterElement0(nums2)
	greaterMap := make(map[int]int, 0)
	for i, v := range nums2 {
		greaterMap[v] = temps[i]
	}

	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = greaterMap[nums1[i]]
	}
	return res
}

func Test_0(t *testing.T) {
	//fmt.Println(nextGreaterElement0([]int{2, 1, 2, 4, 3}))
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}
