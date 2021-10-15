package easy_test

import (
	"fmt"
	"testing"
)

func Test_peakIndexInMountainArray2(t *testing.T) {
	//fmt.Println(peakIndexInMountainArray2([]int{3, 5, 3, 2, 0}))
	fmt.Println(peakIndexInMountainArray2([]int{0, 1, 0}))
}

func peakIndexInMountainArray2(arr []int) int {
	l, r := 1, len(arr)-2
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid-1] < arr[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
