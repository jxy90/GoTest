package easy_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_kWeakestRows(t *testing.T) {
	fmt.Println(kWeakestRows([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1}}, 3))
}

//朴素法,mat行求和,比较排序
func kWeakestRows(mat [][]int, k int) []int {
	m := len(mat)
	//n := len(mat[0])
	//定义array和mat的行对应.
	array := make([]int, m)
	index := make([]int, m)
	for i := range mat {
		count := 0
		for j := range mat[i] {
			count += mat[i][j]
		}
		array[i] = count
		index[i] = i
	}
	sort.Slice(index, func(i, j int) bool {
		a, b := array[index[i]], array[index[j]]
		return a < b || (a == b && index[i] < index[j])
	})
	return index[:k]
}
