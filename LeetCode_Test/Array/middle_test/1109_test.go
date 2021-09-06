package middle_test

import (
	"fmt"
	"testing"
)

func Test_corpFlightBookings(t *testing.T) {
	fmt.Println(corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5))
}

//差分
func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n+2)
	for _, booking := range bookings {
		nums[booking[0]] += booking[2]
		nums[booking[1]+1] -= booking[2]
	}
	for i := 1; i < n+1; i++ {
		nums[i] += nums[i-1]
	}
	return nums[1 : len(nums)-1]
}
