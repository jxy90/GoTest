package hard_test

import (
	"fmt"
	"testing"
)

func Test_gridIllumination(t *testing.T) {
	fmt.Println(gridIllumination(5, [][]int{{0, 0}, {5, 5}}, [][]int{{1, 1}, {1, 0}}))
}
func gridIllumination(n int, lamps, queries [][]int) []int {
	type pair struct{ x, y int }
	//点是否被访问
	points := map[pair]bool{}
	//行，列，对角线所存在的灯泡数量
	row := map[int]int{}
	col := map[int]int{}
	diagonal := map[int]int{}
	antiDiagonal := map[int]int{}
	for _, lamp := range lamps {
		r, c := lamp[0], lamp[1]
		p := pair{r, c}
		if points[p] {
			continue
		}
		points[p] = true
		row[r]++
		col[c]++
		diagonal[r-c]++
		antiDiagonal[r+c]++
	}

	ans := make([]int, len(queries))
	for i, query := range queries {
		r, c := query[0], query[1]
		//行，列，对角线所存在的亮的灯泡
		if row[r] > 0 || col[c] > 0 || diagonal[r-c] > 0 || antiDiagonal[r+c] > 0 {
			ans[i] = 1
		}
		//九宫格存在亮的灯泡，行，列，对角线减去，对应的值
		for x := r - 1; x <= r+1; x++ {
			for y := c - 1; y <= c+1; y++ {
				if x < 0 || y < 0 || x >= n || y >= n || !points[pair{x, y}] {
					continue
				}
				delete(points, pair{x, y})
				row[x]--
				col[y]--
				diagonal[x-y]--
				antiDiagonal[x+y]--
			}
		}
	}
	return ans
}
