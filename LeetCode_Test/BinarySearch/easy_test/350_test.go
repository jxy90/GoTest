package easy_test

import (
	"sort"
	"testing"
)

func Test_intersect(t *testing.T) {
	println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	p1, p2 := 0, 0
	n := len(nums1)
	m := len(nums2)
	ans := make([]int, 0)
	for p1 < n && p2 < m {
		if nums1[p1] == nums2[p2] {
			ans = append(ans, nums1[p1])
			p1++
			p2++
		} else if nums1[p1] > nums2[p2] {
			p2++
		} else {
			p1++
		}
	}
	return ans
}
