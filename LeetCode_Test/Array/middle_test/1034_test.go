package middle_test

import (
	"fmt"
	"testing"
)

func Test_colorBorder(t *testing.T) {
	fmt.Println(colorBorder([][]int{{1, 1}, {1, 2}}, 0, 0, 3))
	fmt.Println(colorBorder([][]int{{1, 2, 2}, {2, 3, 2}}, 0, 1, 3))
	fmt.Println(colorBorder([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 1, 1, 2))
	fmt.Println(colorBorder([][]int{{1, 2, 1, 2, 1, 2}, {2, 2, 2, 2, 1, 2}, {1, 2, 2, 2, 1, 2}}, 1, 3, 1))
}

type Point struct {
	x, y int
}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	startColor := grid[row][col]
	options := []int{1, 0, -1, 0, 1}
	visited := map[int]bool{}
	changedPoints := make([]Point, 0)
	//BFS
	queue := make([]Point, 0, m*n)
	queue = append(queue, Point{row, col})
	visited[row*n+col] = true
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		if check(grid, cur, startColor, 0) {
			grid[cur.x][cur.y] = 0
			changedPoints = append(changedPoints, cur)
		}
		for i := 0; i < 4; i++ {
			next := Point{
				x: cur.x + options[i],
				y: cur.y + options[i+1],
			}
			if 0 <= next.x && next.x < m && 0 <= next.y && next.y < n && startColor == grid[next.x][next.y] && !visited[next.x*n+next.y] {
				visited[next.x*n+next.y] = true
				queue = append(queue, next)
				//fmt.Println(next)
			}
		}
	}
	for _, point := range changedPoints {
		grid[point.x][point.y] = color
	}
	return grid
}

func check(grid [][]int, cur Point, startColor, targetColor int) bool {
	//m, n := len(grid), len(grid[0])
	options := []int{1, 0, -1, 0, 1}
	if cur.x == 0 || cur.y == 0 || cur.x == len(grid)-1 || cur.y == len(grid[0])-1 {
		return true
	}
	for i := 0; i < 4; i++ {
		next := Point{
			x: cur.x + options[i],
			y: cur.y + options[i+1],
		}
		if grid[next.x][next.y] != startColor && grid[next.x][next.y] != targetColor {
			return true
		}
	}
	return false
}
