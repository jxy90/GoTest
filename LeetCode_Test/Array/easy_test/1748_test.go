package easy_test

import (
	"fmt"
	"testing"
)

func Test_sumOfUnique(t *testing.T) {
	fmt.Println(sumOfUnique([]int{1, 2, 3, 2}))
	fmt.Println(sumOfUnique([]int{1, 1, 1, 1, 1}))
}

func sumOfUnique(nums []int) (ans int) {
	cache := map[int]int{}
	for _, num := range nums {
		cache[num]++
	}
	for k, v := range cache {
		if v == 1 {
			ans += k
		}
	}
	return
}
