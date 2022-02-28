package middle_test

import (
	"fmt"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	//fmt.Println(singleNumber2([]int{2, 2, 3, 2}))
	fmt.Println(singleNumber([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
	fmt.Println(singleNumber2([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
}

func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		temp := int32(0)
		for _, num := range nums {
			temp += int32(num) >> i & 1
		}
		if temp%3 > 0 {
			ans |= temp % 3 << i
		}
	}
	return int(ans)
}

func singleNumber2(nums []int) int {
	count := [32]int32{}
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			count[i] += int32(num) >> i & 1
		}
	}
	ans := int32(0)
	for i := 0; i < 32; i++ {
		if (count[i]%3)&1 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}
func singleNumber1(nums []int) int {
	count := map[int]int{}
	for _, i := range nums {
		if _, ok := count[i]; ok {
			count[i] += 1
		} else {
			count[i] = 1
		}
	}
	for key, value := range count {
		if value == 1 {
			return key
		}
	}
	return -1
}
