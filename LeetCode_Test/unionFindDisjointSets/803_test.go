package unionFindDisjointSets

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_hitBricks(t *testing.T) {
	//fmt.Println(hitBricks([][]int{{1, 0, 0, 0}, {1, 1, 1, 0}}, [][]int{{1, 0}}))
	//fmt.Println(hitBricks([][]int{{1, 0, 0, 0}, {1, 1, 0, 0}}, [][]int{{1, 1}, {1, 0}}))
	//fmt.Println(hitBricks([][]int{{1}, {1}, {1}, {1}, {1}}, [][]int{{3, 0}, {4, 0}, {1, 0}, {2, 0}, {0, 0}}))
	//fmt.Println(hitBricks([][]int{{1, 0, 1}, {1, 1, 1}}, [][]int{{0, 0}, {0, 2}, {1, 1}}))
	//fmt.Println(hitBricks([][]int{{0, 1, 1, 1, 1}, {1, 1, 1, 1, 0}, {1, 1, 1, 1, 0}, {0, 0, 1, 1, 0}, {0, 0, 1, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
	//	[][]int{{6, 0}, {1, 0}}))
	fmt.Println(hitBricks([][]int{{1, 1, 1, 0, 1, 1, 1, 1}, {1, 0, 0, 0, 0, 1, 1, 1}, {1, 1, 1, 0, 0, 0, 1, 1}, {1, 1, 0, 0, 0, 0, 0, 0}, {1, 0, 0, 0, 0, 0, 0, 0}, {1, 0, 0, 0, 0, 0, 0, 0}},
		[][]int{{4, 6}, {3, 0}}))
}

func hitBricks(grid [][]int, hits [][]int) []int {
	m := len(grid)
	n := len(grid[0])
	//1.去除被打击块
	c := make([][]int, len(grid))
	for i := range c {
		c[i] = make([]int, len(grid[0]))
		for j := range grid[i] {
			c[i][j] = grid[i][j]
		}
	}
	for i := range hits {
		hitX := hits[i][0]
		hitY := hits[i][1]
		c[hitX][hitY] = 0
	}
	//2. 建立union
	//设置根节点
	root := m * n
	u := ConstructorUnionFind803(root + 1)
	//顶部链接根节点
	for i := 0; i < n; i++ {
		if c[0][i] == 1 {
			u.union(i, root)
		}
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if c[i][j] == 1 {
				if i == 2 {
					fmt.Println(1)
				}
				if c[i-1][j] == 1 {
					//上
					u.union((i-1)*n+j, i*n+j)
				}
				if j-1 >= 0 && c[i][j-1] == 1 {
					//左
					u.union(i*n+j-1, i*n+j)
				}
			}
		}
	}
	//3. 反向添加hits
	hitslen := len(hits)
	ans := make([]int, hitslen)
	for i := hitslen - 1; i >= 0; i-- {
		hitX := hits[i][0]
		hitY := hits[i][1]
		if grid[hitX][hitY] == 0 {
			ans[i] = 0
			continue
		}
		before := u.count[u.find(root)]
		if hitX == 0 {
			u.union(hitY, root)
		}
		for j := 0; j < 4; j++ {
			newX := hitX + optionsX[j]
			newY := hitY + optionsY[j]
			if newX < 0 || newX >= m || newY < 0 || newY >= n || c[newX][newY] == 0 {
				continue
			}
			u.union(hitX*n+hitY, newX*n+newY)
		}
		after := u.count[u.find(root)]
		ans[i] = CommonUtil.Max(0, after-before-1)

		c[hitX][hitY] = 1
	}

	return ans
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
	c[total-1] = 0
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
	if u.count[yp] > u.count[xp] {
		u.parent[xp] = yp
		u.count[yp] += u.count[xp]
	} else {
		u.parent[yp] = xp
		u.count[xp] += u.count[yp]
	}
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
