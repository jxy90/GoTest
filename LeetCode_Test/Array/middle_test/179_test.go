package middle_test

import (
	"strconv"
	"strings"
	"testing"
)

func largestNumber(nums []int) string {
	largestNumberSort(nums, 0, len(nums)-1)
	if nums[0] == 0 {
		return "0"
	}
	sb := strings.Builder{}
	for i := range nums {
		sb.WriteString(strconv.Itoa(nums[i]))
	}
	return sb.String()
}

func largestNumberSort(nums []int, start, end int) {
	//if start < end {
	left, right := start, end
	mid := nums[left+(right-left)/2]
	for left <= right {
		for largestNumberCompare(mid, nums[left]) {
			left++
		}
		for largestNumberCompare(nums[right], mid) {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	if start < right {
		largestNumberSort(nums, start, right)
	}
	if left < end {
		largestNumberSort(nums, left, end)
	}
	//}
}

func largestNumberCompare(a, b int) bool {
	as, bs := 10, 10
	for a > as {
		as *= 10
	}
	for b > bs {
		bs *= 10
	}
	return bs*a+b < as*b+a
}

func Test_largestNumber(t *testing.T) {
	nums := []int{3, 30, 34, 5, 9}
	println(largestNumber(nums))
}
