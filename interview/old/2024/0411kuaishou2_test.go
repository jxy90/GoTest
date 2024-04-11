package _024

import (
	"fmt"
	"testing"
)

func Test_JO(t *testing.T) {
	fmt.Printf("Hello, World!")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	helper(nums)
	fmt.Println(nums)
}

//奇偶混杂的数组，最后返回奇数在前,并且保证顺序不变
func helper(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if isX(nums[i]) && !isX(nums[j]) {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

//奇偶混杂的数组，最后返回奇数在前
//func helper(nums []int) {
//	left, right := 0, len(nums)-1
//	for left < right {
//		for !isX(nums[left]) {
//			left++
//		}
//		for isX(nums[right]) {
//			right--
//		}
//		if left < right {
//			nums[left], nums[right] = nums[right], nums[left]
//			left++
//			right--
//		}
//	}
//}

func isX(num int) bool {
	return num%2 == 0
}
