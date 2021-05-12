package unionFindDisjointSets

import (
	"github.com/jxy90/GoTest/LeetCode_Test/Struct"
	"testing"
)

func Test_solve(t *testing.T) {
	//solve([][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}})
	//solve([][]byte{{'O', 'O'}, {'O', 'O'}})
	//solve([][]byte{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}})
	//solve([][]byte{{'X', 'O', 'X'}, {'X', 'O', 'X'}, {'X', 'O', 'X'}})
	solve([][]byte{{'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}, {'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}})
}

//并查集
func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	unionFind := Struct.ConstructorUnionFind(m*n + 1)
	flag := m * n
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				if (i == 0 || j == 0 || i == m-1 || j == n-1) && board[i][j] == 'O' {
					//并查集
					unionFind.Union(i*n+j, flag)
				} else {
					for k := 0; k < 4; k++ {
						newX, newY := i+optionsX[k], j+optionsY[k]
						if 0 <= newX && newX < m && 0 <= newY && newY < n && board[newX][newY] == 'O' {
							unionFind.Union(i*n+j, newX*n+newY)
						}
					}
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !unionFind.Same(i*m+j, flag) {
				board[i][j] = 'X'
			}
		}
	}
	return
}

var optionsX = []int{1, 0, -1, 0}
var optionsY = []int{0, 1, 0, -1}

func solve0(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				if board[i][j] == 'O' {
					//边界情况
					//DFS 记录位置
					solveDFS(board, i, j)
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
	return
}
func solveDFS(board [][]byte, x, y int) {
	if 0 <= x && x < len(board) && 0 <= y && y < len(board[0]) && board[x][y] == 'O' {
		board[x][y] = 'A'
	} else {
		return
	}
	for i := 0; i < 4; i++ {
		newX, newY := x+optionsX[i], y+optionsY[i]
		solveDFS(board, newX, newY)
	}
}
