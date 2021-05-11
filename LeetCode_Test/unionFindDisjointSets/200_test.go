package unionFindDisjointSets

import "testing"

func Test_numIslands(t *testing.T) {
	println(numIslands([][]byte{{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}))
	println(numIslands([][]byte{{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}))
}

//并查集
func numIslands(grid [][]byte) int {

	return 0
}

//BFS
func numIslands0(grid [][]byte) int {
	optionX := []int{0, 1, 0, -1}
	optionY := []int{1, 0, -1, 0}
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	count := 0
	queue := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				queue = append(queue, []int{i, j})
				for len(queue) > 0 {
					now := queue[0]
					grid[now[0]][now[1]] = '0'
					queue = queue[1:]
					for k := 0; k < 4; k++ {
						new := []int{now[0] + optionX[k], now[1] + optionY[k]}
						if 0 <= new[0] && new[0] < m && 0 <= new[1] && new[1] < n && grid[new[0]][new[1]] == '1' {
							queue = append(queue, new)
						}
					}
				}
			}
		}
	}
	return count
}
