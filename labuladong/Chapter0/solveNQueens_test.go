package Chapter0_test

import "testing"

func solveNQueens(n int) [][]string {
	solveNQueensRes = make([][]string, 0, 0)
	board := make([]string, n, n)
	temp := ""
	for i := 0; i < n; i++ {
		temp += "."
	}
	for k := range board {
		board[k] = temp
	}
	solveNQueensBackTrack(board, 0)

	return solveNQueensRes
}

var solveNQueensRes [][]string

func solveNQueensBackTrack(board []string, row int) {
	if row == len(board) {
		temp := make([]string, len(board))
		copy(temp, board)
		solveNQueensRes = append(solveNQueensRes, temp)
		return
	}
	for col := 0; col < len(board[row]); col++ {
		if !solveNQueensValidate(board, row, col) {
			continue
		}
		board[row] = board[row][:col] + "Q" + board[row][col+1:]
		solveNQueensBackTrack(board, row+1)
		board[row] = board[row][:col] + "." + board[row][col+1:]
	}
}

func solveNQueensValidate(board []string, row, col int) bool {
	//lie
	for i := 0; i < row; i++ {
		if string(board[i][col]) == "Q" {
			return false
		}
	}
	//zuoshan
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if string(board[i][j]) == "Q" {
			return false
		}
	}
	//youshang
	for i, j := row-1, col+1; i >= 0 && j < len(board[row]); i, j = i-1, j+1 {
		if string(board[i][j]) == "Q" {
			return false
		}
	}
	return true
}

func Test_solveNQueens(t *testing.T) {
	data := solveNQueens(8)
	println(len(data))
}
