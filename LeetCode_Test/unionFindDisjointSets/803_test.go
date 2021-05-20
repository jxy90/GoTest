package unionFindDisjointSets

import (
	"testing"
)

func Test_hitBricks(t *testing.T) {
	println(hitBricks([][]int{{1, 0, 0, 0}, {1, 1, 1, 0}}, [][]int{{1, 0}}))
}

func hitBricks(grid [][]int, hits [][]int) []int {
	//1.去除被打击块
	copy := grid
	for i := range hits {
		copy[hits[i][0]][hits[i][1]] = 0
	}
	//2. 建立union
	m := len(grid)
	n := len(grid[0])
	//设置根节点
	root := m * n
	u := ConstructorUnionFind803(root + 1)
	//顶部链接根节点
	for i := 0; i < n; i++ {
		if copy[0][i] == 1 {
			u.union(i, root)
		}
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if copy[i][j] == 1 {
				if copy[i-1][j] == 1 {
					//上
					u.union((i-i)*n+j, i*n+j)
				} else if j-1 >= 0 && copy[i][j-1] == 1 {
					//左
					u.union(i*n+j-1, i*n+j)
				}
			}
		}
	}
	//3. 反向添加hits

	return nil
}

type UnionFind803 struct {
	parent []int
	count  []int
}

func ConstructorUnionFind803(total int) *UnionFind803 {
	p := make([]int, total)
	c := make([]int, total)
	for i := 0; i < total; i++ {
		p[i] = i
		c[i] = 1
	}
	return &UnionFind803{
		parent: p,
		count:  c,
	}
}

func (u *UnionFind803) union(x, y int) {
	xp := u.find(x)
	yp := u.find(y)
	if xp == yp {
		return
	}
	u.parent[xp] = yp
	u.count[yp] += u.count[xp]
}
func (u *UnionFind803) find(x int) int {
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
