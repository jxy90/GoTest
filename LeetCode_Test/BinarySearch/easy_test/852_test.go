package easy_test

import (
	"fmt"
	"testing"
)

func Test_peakIndexInMountainArray(t *testing.T) {
	fmt.Println(peakIndexInMountainArray([]int{3, 5, 3, 2, 0}))
	fmt.Println(peakIndexInMountainArray([]int{0, 1, 0}))
}
func peakIndexInMountainArray(arr []int) int {
	m, n := 1, len(arr)-2
	for m <= n {
		mid := (m + n) / 2
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid] < arr[mid-1] {
			n = mid - 1
		} else {
			m = mid + 1
		}
	}
	return m
}
