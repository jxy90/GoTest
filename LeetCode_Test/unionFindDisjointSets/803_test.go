package unionFindDisjointSets

import (
	"testing"
)

func Test_hitBricks(t *testing.T) {
	println(hitBricks([][]int{{1, 0, 0, 0}, {1, 1, 1, 0}}, [][]int{{1, 0}}))
}

func hitBricks(grid [][]int, hits [][]int) []int {
	return nil
}

type UnionFind803 struct {
	parent []int
}
