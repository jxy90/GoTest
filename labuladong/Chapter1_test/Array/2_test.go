package Array

import (
	"testing"
)

//小而美的算法技巧：差分数组

//区间加法
//https://leetcode.cn/problems/range-addition/description/
func getModifiedArray(length int, updates [][]int) []int {
	diffs := make([]int, length)
	nums := make([]int, length)
	for _, update := range updates {
		diffs[update[0]] += update[2]
		if update[1]+1 < length {
			diffs[update[1]+1] -= update[2]
		}
	}
	nums[0] = diffs[0]
	for i := 1; i < length; i++ {
		nums[i] = nums[i-1] + diffs[i]
	}
	return nums
}

//1109. 航班预订统计
//https://leetcode.cn/problems/corporate-flight-bookings/description/
func corpFlightBookings(bookings [][]int, n int) []int {
	diffs := make([]int, n)
	ans := make([]int, n)
	for _, booking := range bookings {
		s, e, v := booking[0]-1, booking[1], booking[2]
		diffs[s] += v
		if e < n {
			diffs[e] -= v
		}
	}
	ans[0] = diffs[0]
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] + diffs[i]
	}
	return ans
}

//拼车
//https://leetcode.cn/problems/car-pooling/
func carPooling(trips [][]int, capacity int) bool {
	n := 1001
	diffs := make([]int, n)
	ans := make([]int, n)
	for _, trip := range trips {
		s, e, v := trip[1], trip[2], trip[0]
		diffs[s] += v
		if e < n {
			diffs[e] -= v
		}
	}
	ans[0] = diffs[0]
	for i := 1; i < n; i++ {
		ans[i] += ans[i-1] + diffs[i]
	}
	for i := range ans {
		if ans[i] > capacity {
			return false
		}
	}
	return true
}

func Test_2(t *testing.T) {
	//fmt.Println(getModifiedArray(5, [][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}}))
}
