package unionFindDisjointSets

import "testing"

func Test_solve(t *testing.T) {
	solve([][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}})
}

func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	//n := len(board[0])
}
