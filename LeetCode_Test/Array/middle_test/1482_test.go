package middle_test

import (
	"fmt"
	"testing"
)

func Test_minDays(t *testing.T) {
	fmt.Println(minDays([]int{1, 10, 3, 10, 2}, 3, 1))
	fmt.Println(minDays([]int{1, 10, 3, 10, 2}, 3, 2))
	fmt.Println(minDays([]int{7, 7, 7, 7, 12, 7, 7}, 2, 3))
	fmt.Println(minDays([]int{1000000000, 1000000000}, 1, 1))
	fmt.Println(minDays([]int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 4, 2))
}

func minDays(bloomDay []int, m int, k int) int {
	n := len(bloomDay)
	if n < m*k {
		return -1
	}
	left := 1
	right := 1000000000
	//for _, value := range bloomDay {
	//	if value < left {
	//		left = value
	//	}
	//	if value > right {
	//		right = value
	//	}
	//}
	for left <= right {
		mid := left + (right-left)/2
		//如果mid满足条件缩小right，反之增大left
		if minDaysValid(bloomDay, m, k, mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func minDaysValid(bloomDay []int, m int, k int, day int) bool {
	block := make([]int, len(bloomDay))
	index := 0
	for i := 0; i < len(bloomDay); i++ {
		if day >= bloomDay[i] {
			block[index]++
		} else {
			index++
		}
	}
	count := 0
	for i := range block {
		count += block[i] / k
	}
	if count >= m {
		return true
	}
	return false
}
