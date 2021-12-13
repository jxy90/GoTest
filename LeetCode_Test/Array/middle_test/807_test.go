package middle_test

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_maxIncreaseKeepingSkyline(t *testing.T) {
	fmt.Println(maxIncreaseKeepingSkyline([][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}))
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	row := make([]int, m)
	col := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if row[i] < grid[i][j] {
				row[i] = grid[i][j]
			}
			if col[j] < grid[i][j] {
				col[j] = grid[i][j]
			}
		}
	}
	ans := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			min := CommonUtil.Min(row[i], col[j])
			if grid[i][j] < min {
				ans += min - grid[i][j]
			}
		}
	}
	return ans
}
