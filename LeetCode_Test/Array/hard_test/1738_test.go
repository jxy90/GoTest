package hard_test

import (
	"testing"
)

func Test_kthLargestValue(t *testing.T) {
	println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 1))
	println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 2))
	println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 3))
	println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 4))
	println(kthLargestValue([][]int{{1, 6, 5, 6, 1}, {2, 8, 2, 10, 2}, {4, 3, 5, 10, 0}, {1, 7, 10, 5, 5}, {4, 6, 9, 4, 1}, {6, 6, 4, 0, 1}, {4, 0, 2, 8, 1}, {8, 2, 4, 9, 7}, {3, 10, 10, 4, 7}, {6, 1, 1, 9, 4}, {9, 0, 2, 6, 6}, {0, 10, 3, 2, 5}, {10, 3, 7, 7, 2}}, 30))
}

func kthLargestValue(matrix [][]int, k int) int {
	k = k - 1
	m := len(matrix)
	n := len(matrix[0])
	f := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		f[i] = make([]int, n+1)
	}
	nums := make([]int, 0)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			f[i][j] = f[i][j-1] ^ f[i-1][j] ^ f[i-1][j-1] ^ matrix[i-1][j-1]
			nums = append(nums, f[i][j])
		}
	}
	kthnums(nums, 0, len(nums)-1, k)
	return nums[k]
	//sort.Ints(nums)
	//return nums[len(nums)-k-1]
}

func kthnums(nums []int, start, end, k int) {
	if start >= end {
		return
	}
	left := start
	right := end
	mid := nums[left+(right-left)/2]
	for left <= right {
		for nums[left] > mid {
			left++
		}
		for nums[right] < mid {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	if k < left {
		kthnums(nums, start, right, k)
	} else {
		kthnums(nums, left, end, k)
	}
}
