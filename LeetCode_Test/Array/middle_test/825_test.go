package middle_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_numFriendRequests(t *testing.T) {
	fmt.Println(numFriendRequests([]int{16, 16}))
	fmt.Println(numFriendRequests([]int{20, 30, 100, 110, 120}))
}

func numFriendRequests(ages []int) int {
	sort.Ints(ages)
	n := len(ages)
	ans := 0
	for i := 0; i < n; i++ {
		begin := helper1(ages, i)
		end := helper2(ages, i)
		ans += end - begin
	}
	return ans
}

func helper1(ages []int, x int) int {
	left, right := 0, x
	for left < right {
		mid := left + (right-left+1)>>1
		midVal := ages[mid]
		if numFriendRequestsCheck(ages[x], midVal) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
func helper2(ages []int, x int) int {
	left, right := x, len(ages)-1
	for left < right {
		mid := left + (right-left+1)>>1
		midVal := ages[mid]
		if numFriendRequestsCheck(ages[x], midVal) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func numFriendRequestsCheck(ageX, ageY int) bool {
	//age[y] <= 0.5 * age[x] + 7
	//age[y] > age[x]
	//age[y] > 100 && age[x] < 100
	if float64(ageY) <= 0.5*float64(ageX)+7 {
		return false
	}
	if ageY > ageX {
		return false
	}
	if ageY > 100 && ageX < 100 {
		return false
	}
	return true
}
