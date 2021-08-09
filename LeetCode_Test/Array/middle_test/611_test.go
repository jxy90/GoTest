package middle_test

import (
	"sort"
	"testing"
)

func Test_triangleNumber(t *testing.T) {
	println(triangleNumber([]int{2, 2, 3, 4}))
}

func triangleNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	ans := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			ans += sort.SearchInts(nums[j+1:], nums[i]+nums[j])
		}
	}
	return ans
}

//TLE
func triangleNumber0(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	ans := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				a, b, c := nums[i], nums[j], nums[k]
				if a+b > c && b-a < c {
					ans++
				}
			}
		}
	}
	return ans
}
