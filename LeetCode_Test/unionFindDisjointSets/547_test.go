package unionFindDisjointSets

import (
	"github.com/jxy90/GoTest/LeetCode_Test/Struct"
	"testing"
)

func Test_findCircleNum(t *testing.T) {
	//println(findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
	println(findCircleNum([][]int{
		{1, 1, 1, 0, 1, 1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 1, 0, 0},
		{1, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0, 1, 0},
		{1, 0, 0, 1, 1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 1, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 1, 0, 0, 1, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
	}))
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	union := Struct.ConstructorUnionFind(n)
	for i := range isConnected {
		for j := range isConnected[i] {
			if isConnected[i][j] == 1 {
				union.Union(i, j)
			}
		}
	}
	memo := make(map[int]bool, 0)
	for i := 0; i < len(union.Parent); i++ {
		if memo[union.Parent[i]] == false {
			memo[union.Find(union.Parent[i])] = true
		}
	}
	return len(memo)
}
