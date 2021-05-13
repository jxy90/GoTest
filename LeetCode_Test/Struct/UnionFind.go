package Struct

type UnionFind struct {
	parent []int
}

func ConstructorUnionFind(total int) UnionFind {
	parent := make([]int, 0)
	for i := 0; i < total; i++ {
		parent = append(parent, i)
	}
	return UnionFind{
		parent: parent,
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
	u.parent[yp] = xp
}

func (u *UnionFind) Same(x, y int) bool {
	xp := u.Find(x)
	yp := u.Find(y)
	return xp == yp
}
