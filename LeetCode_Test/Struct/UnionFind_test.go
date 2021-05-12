package Struct

type UnionFind struct {
	count  int
	parent []int
	rank   []int
}

func ConstructorUnionFind(grid [][]byte) UnionFind {
	//m := len(grid)
	n := len(grid[0])
	count := 0
	parent := make([]int, 0)
	rank := make([]int, 0)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				parent = append(parent, i*n+j)
				count++
			} else {
				parent = append(parent, -1)
			}
			rank = append(rank, 0)
		}
	}
	return UnionFind{
		count:  count,
		parent: parent,
		rank:   rank,
	}
}

func (u *UnionFind) Find(index int) int {
	if u.parent[index] == -1 {
		return -1
	}
	root := index
	for root != u.parent[root] {
		root = u.parent[root]
	}
	//路径压缩
	i, j := index, 0
	for root != u.parent[i] {
		j = u.parent[i]
		u.parent[i] = root
		i = j
	}
	return root
}

func (u *UnionFind) Union(x, y int) {
	xp := u.Find(x)
	yp := u.Find(y)
	if xp == yp {
		return
	}
	u.parent[xp] = yp
}

func (u *UnionFind) Same(x, y int) bool {
	xp := u.Find(x)
	yp := u.Find(y)
	return xp == yp
}
