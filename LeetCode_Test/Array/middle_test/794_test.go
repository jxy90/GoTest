package middle_test

import (
	"fmt"
	"testing"
)

func Test_validTicTacToe(t *testing.T) {
	fmt.Println(validTicTacToe([]string{"O  ", "   ", "   "}))
	fmt.Println(validTicTacToe([]string{"XXX", "   ", "OOO"}))
	fmt.Println(validTicTacToe([]string{"XXX", "XOO", "OO "}))
}
func validTicTacToe(board []string) bool {
	//统计XO的数量
	xCount, oCount := 0, 0
	for _, s := range board {
		for _, b := range s {
			switch b {
			case 'O':
				oCount++
			case 'X':
				xCount++
			}
		}
	}
	a, b := validTicTacToeWin(board, 'X'), validTicTacToeWin(board, 'O')
	//o大于x,失败
	if oCount > xCount || xCount-oCount > 1 {
		return false
	}
	//x赢,o大于等于x,失败
	if a && oCount >= xCount {
		return false
	}
	//o赢,x不等于o失败
	if b && oCount != xCount {
		return false
	}
	//o和x都赢,失败
	if a && b {
		return false
	}
	return true
}

func validTicTacToeWin(board []string, p byte) bool {
	for i := 0; i < 3; i++ {
		if (board[i][0] == p && board[i][1] == p && board[i][2] == p) ||
			(board[0][i] == p && board[1][i] == p && board[2][i] == p) {
			return true
		}
	}
	return (board[0][0] == p && board[1][1] == p && board[2][2] == p) ||
		(board[0][2] == p && board[1][1] == p && board[2][0] == p)
}
