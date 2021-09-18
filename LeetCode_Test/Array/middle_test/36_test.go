package middle_test

import (
	"fmt"
	"testing"
)

func Test_isValidSudoku(t *testing.T) {
	fmt.Println(isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}))
}

func isValidSudoku(board [][]byte) bool {
	cols := make([]map[byte]bool, 9)
	for i := range cols {
		cols[i] = map[byte]bool{}
	}
	rows := make([]map[byte]bool, 9)
	for i := range cols {
		rows[i] = map[byte]bool{}
	}
	tables := make([]map[byte]bool, 9)
	for i := range cols {
		tables[i] = map[byte]bool{}
	}
	for i, bytes := range board {
		for j, b := range bytes {
			//行
			if isValidSudokuContains(rows[i], b) {
				return false
			}
			//列
			if isValidSudokuContains(cols[j], b) {
				return false
			}
			//格
			k := i/3*3 + j/3
			if isValidSudokuContains(tables[k], b) {
				return false
			}
		}
	}
	return true
}

func isValidSudokuContains(cache map[byte]bool, b byte) bool {
	if b == '.' {
		return false
	}
	if cache[b] {
		return true
	}
	cache[b] = true
	return false
}
