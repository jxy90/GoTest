package easy_test

import (
	"fmt"
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix)-1, len(matrix[0])-1
	for i := 0; i <= m; i++ {
		l, r := 0, n
		for l < r {
			mid := (l + r + 1) >> 1
			if matrix[i][mid] == target {
				return true
			} else if matrix[i][mid] > target {
				r = mid - 1
			} else {
				l = mid
			}
		}
		if matrix[i][l] == target {
			return true
		}
	}
	return false
}
