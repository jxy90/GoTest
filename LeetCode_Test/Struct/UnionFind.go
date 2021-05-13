package Struct

type UnionFind struct {
	Parent []int
}

func ConstructorUnionFind(total int) UnionFind {
	parent := make([]int, 0)
	for i := 0; i < total; i++ {
		parent = append(parent, i)
	}
	return UnionFind{
		Parent: parent,
	}
}

func (u *UnionFind) Find(index int) int {
	if u.Parent[index] == -1 {
		return -1
	}
	root := index
	for root != u.Parent[root] {
		root = u.Parent[root]
	}
	//路径压缩
	i, j := index, 0
	for root != u.Parent[i] {
		j = u.Parent[i]
		u.Parent[i] = root
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
	u.Parent[yp] = xp
}

func (u *UnionFind) Same(x, y int) bool {
	xp := u.Find(x)
	yp := u.Find(y)
	return xp == yp
}
