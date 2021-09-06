package middle_test_test

import (
	"fmt"
	"testing"
)

func Test_xorQueries(t *testing.T) {
	fmt.Println(xorQueries([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}))
}

//前缀和
func xorQueries(arr []int, queries [][]int) []int {
	ans := make([]int, 0)
	n := len(arr)
	f := make([]int, n)
	for i := range f {
		if i == 0 {
			f[i] = arr[i]
		} else {
			f[i] = f[i-1] ^ arr[i]
		}
	}

	for i := range queries {
		l, r := queries[i][0]-1, queries[i][1]
		lv, rv := 0, f[r]
		if l == -1 {
			lv = 0
		} else {
			lv = f[l]
		}
		ans = append(ans, lv^rv)
	}
	return ans
}

//暴力方法
func xorQueries0(arr []int, queries [][]int) []int {
	ans := make([]int, 0)
	for i := range queries {
		temp := 0
		for j := queries[i][0]; j <= queries[i][1]; j++ {
			temp ^= arr[j]
		}
		ans = append(ans, temp)
	}
	return ans
}
