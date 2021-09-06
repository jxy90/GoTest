package easy_test

import (
	"fmt"
	"testing"
)

func Test_majorityElement2(t *testing.T) {
	fmt.Println(majorityElement([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
}

func majorityElement2(nums []int) int {
	cache := make(map[int]int)
	ans := 0
	for _, num := range nums {
		cache[num]++
		if cache[num] > cache[ans] {
			ans = num
		}
	}
	if cache[ans]*2 < len(nums) {
		return -1
	}
	return ans
}
