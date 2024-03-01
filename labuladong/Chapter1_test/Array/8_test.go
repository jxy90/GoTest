package Array

import (
	"fmt"
	"sort"
	"testing"
)

//田忌赛马背后的算法决策

type Pair struct {
	i, j int
}

//优势洗牌
func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	pairs := make([]Pair, 0, n)
	for i, i2 := range nums2 {
		pairs = append(pairs, Pair{i, i2})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].j < pairs[j].j
	})
	sort.Ints(nums1)
	left, right := 0, n-1
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		index, val := pairs[i].i, pairs[i].j
		if nums1[right] > val {
			ans[index] = nums1[right]
			right--
		} else {
			ans[index] = nums1[left]
			left++
		}
	}
	return ans
}

func Test_8(t *testing.T) {
	fmt.Println(advantageCount([]int{2, 7, 11, 15}, []int{1, 10, 4, 11}))
}
