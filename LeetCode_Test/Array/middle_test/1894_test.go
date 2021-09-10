package middle_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_chalkReplacer(t *testing.T) {
	fmt.Println(chalkReplacer([]int{5, 1, 5}, 22))
	fmt.Println(chalkReplacer([]int{3, 4, 1, 2}, 25))
}

func chalkReplacer(chalk []int, k int) int {
	n := len(chalk)
	sum := make([]int, n+1)
	for i, i2 := range chalk {
		sum[i+1] = sum[i] + i2
	}
	if sum[n] <= k {
		k = k % sum[n]
	}
	sum = sum[1:]
	left, right := 0, n-1
	for left < right {
		mid := (left + right) >> 1
		if sum[mid] < k {
			left = mid + 1
		} else if sum[mid] == k {
			return mid + 1
		} else {
			right = mid
		}
	}
	if sum[left] == k {
		return left + 1
	}
	return left
}

func chalkReplacer0(chalk []int, k int) int {
	n := len(chalk)
	sum := make([]int, n+1)
	for i, i2 := range chalk {
		sum[i+1] = sum[i] + i2
	}
	if sum[n] <= k {
		k = k % sum[n]
	}
	index := sort.SearchInts(sum[1:], k)
	if sum[index] == k {
		return index + 1
	}
	return index
}
