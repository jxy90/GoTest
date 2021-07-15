package middle_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"sort"
	"testing"
)

func Test_maximumElementAfterDecrementingAndRearranging(t *testing.T) {
	println(maximumElementAfterDecrementingAndRearranging([]int{2, 2, 1, 2, 1}))
	println(maximumElementAfterDecrementingAndRearranging([]int{100, 1, 1000}))
}

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	n := len(arr)
	count := make([]int, n+1)
	for _, v := range arr {
		count[CommonUtil.Min(v, n)]++
	}
	//从1开始计算,包含第一个数字
	cnt := 0
	for _, v := range count[1:] {
		if v == 0 {
			cnt++
		} else {
			cnt -= CommonUtil.Min(v-1, cnt)
		}
	}
	return n - cnt
}

func maximumElementAfterDecrementingAndRearranging0(arr []int) int {
	n := len(arr)
	sort.Ints(arr)
	arr[0] = 1
	for i := 1; i < n; i++ {
		arr[i] = CommonUtil.Min(arr[i], arr[i-1]+1)
	}
	return arr[n-1]
}
