package old

import (
	"fmt"
	"testing"
)

func Test_循环智能(t *testing.T) {
	fmt.Println(FindKthLargeNumber([]int{1, 3, 5, 6, 7, 8, 9, 4, 2}, 2))
	fmt.Println(QuickSort([]int{1, 3, 5, 6, 7, 8, 9, 4, 2}, 0, 8))
}

//2.无序整数数组
func FindKthLargeNumber(data []int, k int) int {
	n := len(data)
	FindKthLargeNumberHelper(data, n-k, 0, n-1)
	//fmt.Println(data)
	//fmt.Println(n - k - 1)
	return data[n-k]
}

func FindKthLargeNumberHelper(data []int, k, start, end int) {
	if start < end {
		left, right := start, end
		mid := data[(left+right)/2]
		for left <= right {
			for data[left] < mid {
				left++
			}
			for data[right] > mid {
				right--
			}
			if left <= right {
				data[left], data[right] = data[right], data[left]
				left++
				right--
			}
		}
		if right > start && k >= start && k <= right {
			FindKthLargeNumberHelper(data, k, start, right)
		}
		if left < end && k >= left && k <= end {
			FindKthLargeNumberHelper(data, k, left, end)
		}
	}
}

//3.反转链表

//4.两个单词最少操作数

//5.快排
func QuickSort(data []int, start, end int) (sortedData []int) {
	if start < end {
		left, right := start, end
		mid := data[(left+right)/2]
		for left <= right {
			for data[left] < mid {
				left++
			}
			for data[right] > mid {
				right--
			}
			if left <= right {
				data[left], data[right] = data[right], data[left]
				left++
				right--
			}
		}
		if right > start {
			QuickSort(data, start, right)
		}
		if left < end {
			QuickSort(data, left, end)
		}
	}
	return data
}
