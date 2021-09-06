package easy_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_fairCandySwap(t *testing.T) {
	fmt.Println(fairCandySwap([]int{35, 17, 4, 24, 10}, []int{63, 21}))
}

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	//sort.Ints(aliceSizes)
	sort.Ints(bobSizes)
	sumA, sumB := 0, 0
	cacheA, cacheB := map[int]int{}, map[int]int{}
	for _, size := range aliceSizes {
		sumA += size
		cacheA[size]++
	}
	for _, size := range bobSizes {
		sumB += size
		cacheB[size]++
	}
	diff := sumA - sumB
	heler := func(target int) bool {
		l, r := 0, len(bobSizes)-1
		for l <= r {
			mid := l + (r-l)>>1
			val := bobSizes[mid]
			if val == target {
				return true
			} else if val > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return false
	}
	for _, size := range aliceSizes {
		if heler(size - diff/2) {
			return []int{size, size - diff/2}
		}
	}
	ans := []int{0, 0}
	return ans
}
