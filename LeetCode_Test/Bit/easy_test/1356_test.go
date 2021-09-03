package easy_test

import (
	"sort"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(sortByBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}))
	fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
}

func sortByBits(arr []int) []int {
	counts := make(map[int]int)
	for i := range arr {
		counts[arr[i]] = sortByBitsCount(arr[i])
	}
	sort.Slice(arr, func(i, j int) bool {
		return counts[arr[i]] < counts[arr[j]] || (counts[arr[i]] == counts[arr[j]] && arr[i] < arr[j])
	})
	return arr
}

func sortByBitsCount(n int) int {
	res := 0
	for n > 0 {
		if n&1 == 1 {
			res++
		}
		n >>= 1
	}

	return res
}
