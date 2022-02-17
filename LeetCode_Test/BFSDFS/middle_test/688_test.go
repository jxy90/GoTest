package middle_test

import (
	"fmt"
	"testing"
)

func Test_knightProbability(t *testing.T) {
	fmt.Println(knightProbability(3, 2, 0, 0))
}

type pair struct {
	x, y int
}

func knightProbability(n int, k int, row int, column int) float64 {
	options := []pair{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {1, 2}, {-1, 2}, {1, -2}, {-1, -2}}
	dp := make([][][]float64, k+1)
	for step := range dp {
		dp[step] = make([][]float64, n)
		for i := range dp[step] {
			dp[step][i] = make([]float64, n)
			for j := range dp[step][i] {
				if step == 0 {
					dp[step][i][j] = 1
				} else {
					for _, option := range options {
						newRow, newColumn := row+option.x, column+option.y
						if 0 <= newRow && newRow < n && 0 <= newColumn && newColumn < n {
							dp[step][i][j] += dp[step-1][newRow][newColumn] / 8
						}
					}
				}
			}
		}
	}
	return dp[k][row][column]
}
