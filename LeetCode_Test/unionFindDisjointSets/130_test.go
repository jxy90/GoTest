package unionFindDisjointSets

import (
	"fmt"
	"testing"
)

func Test_solve(t *testing.T) {
	solve([][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}})
}

func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	points := make(map[string]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || j == 0 || i == m-1 || j == n-1) && board[i][j] == 'O' {
				//边界情况
				//DFS 记录位置
				solveDFS(board, i, j, &points)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if points[fmt.Sprint(i)+"_"+fmt.Sprint(j)] != true && board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

var optionsX = []int{1, 0, -1, 0}
var optionsY = []int{0, 1, 0, -1}

func solveDFS(board [][]byte, x, y int, points *map[string]bool) {
	(*points)[fmt.Sprint(x)+"_"+fmt.Sprint(y)] = true
	for i := 0; i < 4; i++ {
		newX, newY := x+optionsX[i], y+optionsY[i]
		if board[newX][newY] == 'O' && 0 <= newX && newX < len(board) && 0 <= newY && newY <= len(board[0]) {
			solveDFS(board, newX, newY, points)
		}
	}
}
