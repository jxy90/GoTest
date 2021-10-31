package middle_test

import (
	"fmt"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
}

func singleNumber(nums []int) []int {
	sum := 0
	for _, num := range nums {
		sum ^= num
	}
	index := 0
	for (sum>>index)&1 != 1 {
		index++
	}
	num1, num2 := 0, 0
	for _, num := range nums {
		if num>>index&1 == 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}
	return []int{num1, num2}
}
