package easy_test

import (
	"testing"
)

func Test_883(t *testing.T) {
	t.Log(projectionArea([][]int{{1, 2}, {3, 4}}))
}

func projectionArea(grid [][]int) int {
	var xyArea, yzArea, zxArea int
	for i, row := range grid {
		yzHeight, zxHeight := 0, 0
		for j, v := range row {
			if v > 0 {
				xyArea++
			}
			yzHeight = max(yzHeight, grid[j][i])
			zxHeight = max(zxHeight, v)
		}
		yzArea += yzHeight
		zxArea += zxHeight
	}
	return xyArea + yzArea + zxArea
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
