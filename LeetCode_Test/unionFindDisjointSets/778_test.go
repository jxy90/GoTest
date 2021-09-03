package unionFindDisjointSets

import (
	"testing"
)

func Test_swimInWater(t *testing.T) {
	//fmt.Println(swimInWater([][]int{{0, 2}, {1, 3}}))
	fmt.Println(swimInWater([][]int{{3, 2}, {1, 0}}))
	//fmt.Println(swimInWater([][]int{{0,1,2,3,4},{24,23,22,21,5},{12,13,14,15,16},{11,17,18,19,20},{10,9,8,7,6}}))
}

func swimInWater(grid [][]int) int {
	n := len(grid)
	total := n * n
	left, right := 0, total
	for left <= right {
		mid := left + (right-left)/2
		if validateMid(grid, total, mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func validateMid(grid [][]int, total, mid int) bool {
	union := Constructor778(total)
	for i := range grid {
		for j := range grid[i] {
			if i == 0 && j == 0 {
				continue
			}
			if grid[i][j] > mid {
				continue
			}
			ni, nj := i, j-1
			//左
			if 0 <= nj && nj < len(grid) && grid[ni][nj] <= mid {
				union.union(getIndex(ni, nj, len(grid)), getIndex(i, j, len(grid)))
			}

			ni, nj = i-1, j
			//上
			if 0 <= ni && ni < len(grid) && grid[ni][nj] <= mid {
				union.union(getIndex(ni, nj, len(grid)), getIndex(i, j, len(grid)))
			}
		}
	}
	return union.find(0) == union.find(total-1)
}

func getIndex(i, j, n int) int {
	return i*n + j
}

type UnionFind778 struct {
	parent []int
}

func Constructor778(total int) *UnionFind778 {
	p := make([]int, total)
	for i := range p {
		p[i] = i
	}
	return &UnionFind778{parent: p}
}

func (u *UnionFind778) union(x, y int) {
	xp := u.find(x)
	yp := u.find(y)
	if xp == yp {
		return
	}
	u.parent[xp] = yp
}

func (u *UnionFind778) find(x int) int {
	root := x
	for root != u.parent[root] {
		root = u.parent[root]
	}
	//compress
	i, j := x, 0
	for root != u.parent[i] {
		j = u.parent[i]
		u.parent[i] = root
		i = j
	}
	return root
}
