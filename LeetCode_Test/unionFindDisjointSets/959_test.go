package unionFindDisjointSets

import (
	"testing"
)

func Test_regionsBySlashes(t *testing.T) {
	fmt.Println(regionsBySlashes([]string{" /", "/ "}))
	fmt.Println(regionsBySlashes([]string{" /", "  "}))
	fmt.Println(regionsBySlashes([]string{"\\/", "/\\"}))
	fmt.Println(regionsBySlashes([]string{"/\\", "\\/"}))
	fmt.Println(regionsBySlashes([]string{"//", "/ "}))
}

func regionsBySlashes(grid []string) int {
	m := len(grid)
	n := len(grid[0])
	union := ConstructorUnionFind959(m * n * 4)
	for i := range grid {
		for j := range grid[i] {
			col := i*n + j
			if grid[i][j] == '/' || grid[i][j] == ' ' {
				union.union(col*4, col*4+1)
				union.union(col*4+2, col*4+3)
			}
			if grid[i][j] == '\\' || grid[i][j] == ' ' {
				union.union(col*4, col*4+3)
				union.union(col*4+1, col*4+2)
			}
			//向右合并
			if j+1 < m {
				union.union(col*4+2, (col+1)*4)
			}
			//向下合并
			if i+1 < n {
				union.union(col*4+3, (col+n)*4+1)
			}
		}
	}
	return union.count
}

type UnionFind959 struct {
	parent []int
	count  int
}

func ConstructorUnionFind959(total int) *UnionFind959 {
	p := make([]int, total)
	for i := range p {
		p[i] = i
	}
	return &UnionFind959{
		parent: p,
		count:  total,
	}
}

func (u *UnionFind959) union(x, y int) {
	xp := u.find(x)
	yp := u.find(y)
	if xp == yp {
		return
	}
	u.parent[xp] = yp
	u.count--
}

func (u *UnionFind959) find(x int) int {
	root := x
	for root != u.parent[root] {
		root = u.parent[root]
	}
	//压缩
	i, j := x, 0
	for root != u.parent[i] {
		j = u.parent[i]
		u.parent[i] = root
		i = j
	}
	return root
}
