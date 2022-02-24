package middle_test

import (
	"fmt"
	"testing"
)

func Test_findBall(t *testing.T) {
	fmt.Println(findBall([][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}))
}

func findBall(grid [][]int) []int {
	n := len(grid[0])
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		//球的位置
		col := i
		for _, row := range grid {
			//移动球
			dir := row[col]
			col += dir
			if col < 0 || col >= n || row[col] != dir {
				col = -1
				break
			}
		}
		ans[i] = col
	}
	return ans
}
