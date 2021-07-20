package easy_test

import (
	"sort"
	"testing"
)

func Test_intersection(t *testing.T) {
	println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
}

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	p1, p2 := 0, 0
	n := len(nums1)
	m := len(nums2)
	cache := map[int]int{}
	for p1 < n && p2 < m {
		if nums1[p1] == nums2[p2] {
			cache[nums1[p1]]++
			p1++
			p2++
		} else if nums1[p1] > nums2[p2] {
			p2++
		} else {
			p1++
		}
	}
	ans := make([]int, 0)
	for key := range cache {
		ans = append(ans, key)
	}
	return ans
}
