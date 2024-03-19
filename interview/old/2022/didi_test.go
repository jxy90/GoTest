package _022

import (
	"fmt"
	"testing"
)

func Test_didi(t *testing.T) {
	fmt.Println(count([]int{1, 2, 2, 3, 4, 5, 6, 4, 3, 3, 2, 1, 1, 1}))
	fmt.Println(count([]int{1, 2, 100, 1}))
	fmt.Println(count([]int{}))
}

//6
//1, 2, 2, 3, 4, 5, 6,4,3,3,2,1,1,1
//o(n) o(1)

func count(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	res := 1
	left, right := 0, n-1
	val := min(nums[left], nums[right])
	for left < right {
		for nums[left] <= val {
			left++
		}
		for nums[right] <= val {
			right--
		}
		val = min(nums[left], nums[right])
		res++
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
