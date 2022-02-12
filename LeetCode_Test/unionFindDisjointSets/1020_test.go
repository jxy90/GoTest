package unionFindDisjointSets

import (
	"fmt"
	"testing"
)

func Test_numEnclaves(t *testing.T) {
	fmt.Println(numEnclaves([][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}))
}

func numEnclaves(grid [][]int) int {
	options := []int{1, 0, -1, 0, 1}
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x > m-1 || y < 0 || y > n-1 || visited[x][y] || grid[x][y] == 0 {
			return
		}
		visited[x][y] = true
		for i := 0; i < 4; i++ {
			dfs(x+options[i], y+options[i+1])
		}
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for i := 0; i < n; i++ {
		dfs(0, i)
		dfs(m-1, i)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				ans++
			}
		}
	}
	return ans
}
