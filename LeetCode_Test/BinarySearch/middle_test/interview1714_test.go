package middle_test

import (
	"fmt"
	"testing"
)

func Test_smallestK(t *testing.T) {
	//fmt.Println(smallestK([]int{1, 3, 5, 7, 2, 4, 6, 8}, 4))
	fmt.Println(smallestK([]int{62577, -220, -8737, -22, -6, 59956, 5363, -16699, 0, -10603, 64, -24528, -4818, 96, 5747, 2638, -223, 37663, -390, 35778, -4977, -3834, -56074, 7, -76, 601, -1712, -48874, 31, 3, -9417, -33152, 775, 9396, 60947, -1919, 683, -37092, -524, -8, 1458, 80, -8, 1, 7, -355, 9, 397, -30, -21019, -565, 8762, -4, 531, -211, -23702, 3, 3399, -67, 64542, 39546, 52500, -6263, 4, -16, -1, 861, 5134, 8, 63701, 40202, 43349, -4283, -3, -22721, -6, 42754, -726, 118, 51, 539, 790, -9972, 41752, 0, 31, -23957, -714, -446, 4, -61087, 84, -140, 6, 53, -48496, 9, -15357, 402, 5541, 4, 53936, 6, 3, 37591, 7, 30, -7197, -26607, 202, 140, -4, -7410, 2031, -715, 4, -60981, 365, -23620, -41, 4, -2482, -59, 5, -911, 52, 50068, 38, 61, 664, 0, -868, 8681, -8, 8, 29, 412}, 131))
}

func smallestK(arr []int, k int) []int {
	length := len(arr)
	if length == 0 {
		return []int{}
	}
	smallestKHelper(arr, k)
	//return arr[:k]
	return arr
}

func smallestKHelper(arr []int, k int) []int {
	if len(arr) < 2 {
		return arr
	}
	head, tail := 0, len(arr)-1
	mid := arr[0]
	i := 1
	for head < tail {
		if arr[i] < mid {
			arr[head], arr[i] = arr[i], arr[head]
			head++
			i++
		} else {
			arr[tail], arr[i] = arr[i], arr[tail]
			tail--
		}
	}
	if i < k {
		smallestKHelper(arr[i:], k)
	} else if i > k {
		smallestKHelper(arr[:i], k)
	}
	return arr
}
