package middle_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"sort"
	"testing"
)

func Test_maxFrequency(t *testing.T) {
	fmt.Println(maxFrequency([]int{1000}, 1000))
	fmt.Println(maxFrequency([]int{1, 2, 4}, 5))
}

func maxFrequency(nums []int, k int) int {

	sort.Ints(nums)
	n := len(nums)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	check := func(length int) bool {
		for l := 0; l+length-1 < n; l++ {
			r := l + length - 1
			cur, target := sum[r+1]-sum[l], nums[r]*length
			if target-cur <= k {
				return true
			}
		}
		return false
	}
	l, r := 0, n
	for l <= r {
		mid := (l + r + 1) >> 1
		if check(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}

func maxFrequency0(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	max := 1
	for i := 0; i < n; i++ {
		target := nums[i]
		tempk := k
		temp_max := 0
		for j := i; j >= 0; j-- {
			if target-nums[j] <= tempk {
				tempk -= target - nums[j]
				temp_max++
			} else {
				break
			}
		}
		max = CommonUtil.Max(max, temp_max)
	}
	return max
}
