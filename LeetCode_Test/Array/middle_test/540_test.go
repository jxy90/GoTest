package middle_test

import (
	"fmt"
	"testing"
)

func Test_singleNonDuplicate(t *testing.T) {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	//[3,3,7,7,10,11,11]
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))

}

func singleNonDuplicate(nums []int) int {
	n := len(nums) - 1
	if n == 0 {
		return nums[0]
	}
	left, right := 0, n
	for left < right {
		mid := left + (right-left)>>1
		//偶数^1=偶数+1
		//奇数^1=奇数-1
		//所以mid和mid^1为每次比较的两个数
		if nums[mid] == nums[mid^1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
func singleNonDuplicate0(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans = ans ^ num
	}
	return ans
}
