package unionFindDisjointSets

import (
	"testing"
)

func Test_minSwapsCouples(t *testing.T) {
	println(minSwapsCouples([]int{0, 2, 1, 3}))
}

func minSwapsCouples(row []int) int {
	n := len(row)
	N := n / 2
	uf := ConstructorUnionFind765(N)
	for i := 0; i < n; i += 2 {
		uf.union(row[i]/2, row[(i+1)]/2)
	}
	return N - uf.count
}

type UnionFind765 struct {
	parent []int
	count  int
}

func ConstructorUnionFind765(total int) *UnionFind765 {
	p := make([]int, total)
	for i := 0; i < total; i++ {
		p[i] = i
	}
	return &UnionFind765{
		parent: p,
		count:  total,
	}
}
func (u *UnionFind765) union(x, y int) {
	xp := u.find(x)
	yp := u.find(y)
	if xp == yp {
		return
	}
	u.parent[xp] = yp
	u.count--
}
func (u *UnionFind765) find(x int) int {
	root := x
	for root != u.parent[root] {
		root = u.parent[root]
	}
	i, j := x, 0
	for root != u.parent[i] {
		j = u.parent[i]
		u.parent[i] = root
		i = j
	}
	return root
}
