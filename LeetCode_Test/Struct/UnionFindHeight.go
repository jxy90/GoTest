package Struct

type UnionFindHeight struct {
	Parent []int
	Height []int
}

func ConstructorUnionFindHeight(total int) *UnionFindHeight {
	parent := make([]int, 0)
	height := make([]int, 0)
	for i := 0; i < total; i++ {
		parent = append(parent, i)
		height = append(height, 1)
	}
	return &UnionFindHeight{
		Parent: parent,
		Height: height,
	}
}

func (u *UnionFindHeight) Find(x int) int {
	root := x
	for root != u.Parent[root] {
		root = u.Parent[root]
	}
	//压缩
	i, j := x, 0
	for root != u.Parent[i] {
		j = u.Parent[i]
		u.Parent[i] = root
		i = j
	}
	return root
}
func (u *UnionFindHeight) Union(x, y int) {
	xp := u.Find(x)
	yp := u.Find(y)
	if xp == yp {
		return
	}
	if u.Height[xp] < u.Height[yp] {
		x, y = y, x
		xp, yp = yp, xp
	}
	u.Parent[y] = x
	u.Height[xp] += u.Height[yp]
}
func (u *UnionFindHeight) Same(x, y int) bool {
	return u.Find(x) == u.Find(y)
}
