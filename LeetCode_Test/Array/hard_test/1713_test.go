package hard_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"sort"
	"testing"
)

func Test_minOperations(t *testing.T) {
	fmt.Println(minOperations([]int{5, 1, 3}, []int{9, 4, 2, 3, 4}))
}

func minOperations(target, arr []int) int {
	n := len(target)
	pos := make(map[int]int, n)
	for i, val := range target {
		pos[val] = i
	}
	d := []int{}
	for _, val := range arr {
		if idx, has := pos[val]; has {
			if p := sort.SearchInts(d, idx); p < len(d) {
				d[p] = idx
			} else {
				d = append(d, idx)
			}
		}
	}
	return n - len(d)
}

//OTL
func minOperations1(target []int, arr []int) int {
	m := len(target)
	//n := len(arr)
	target1 := make(map[int]int, m)
	arr1 := make([]int, 0)
	for i := range target {
		target1[target[i]] = i
	}
	for _, val := range arr {
		if _, ok := target1[val]; ok {
			arr1 = append(arr1, target1[val])
		}
	}
	n := len(arr1)
	f := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		f[i] = 1
		for j := 1; j < i; j++ {
			if arr1[i-1] > arr1[j-1] {
				f[i] = CommonUtil.Max(f[i], f[j]+1)
			}
		}
	}
	max := 0
	for i := range f {
		max = CommonUtil.Max(max, f[i])
	}
	return len(target) - max
}

//oom
func minOperations0(target []int, arr []int) int {
	m := len(target)
	n := len(arr)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if target[i-1] == arr[j-1] {
				f[i][j] = f[i-1][j-1] + 1
			} else {
				f[i][j] = CommonUtil.Max(f[i-1][j], f[i][j-1])
			}
		}
	}
	return m - f[m][n]
}
