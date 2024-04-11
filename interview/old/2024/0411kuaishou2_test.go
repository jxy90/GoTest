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

//1 7 3 9 5 2 6 4 8
// func helper(nums []int){
//     cnt := 0
//     for _,v :=range nums{
//         if !isX(v) {
//             cnt++
//         }
//     }
//     p1,p2,p3:=0,0,cnt
//     for p1<len(nums) && p2 <len(nums){
//         for !isX(nums[p1]){
//             p1++
//         }
//         for isX(nums[p2]){
//             p2++
//         }
//     }

//     for p1<cnt && p2<len(nums) {
//         for !isX(nums[p1]){
//             p1++
//         }
//         if p1 < cnt {
//             nums[p1],nums[p2]=nums[p2],nums[p1]
//             p2++
//         }
//     }
// }

//奇偶混杂的数组，最后返回奇数在前
func helper(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		for !isX(nums[left]) {
			left++
		}
		for isX(nums[right]) {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
}

func isX(num int) bool {
	return num%2 == 0
}
