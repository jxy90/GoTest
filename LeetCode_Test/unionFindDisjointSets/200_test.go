package unionFindDisjointSets

import (
	"testing"
)

func Test_numIslands(t *testing.T) {
	//fmt.Println(numIslands([][]byte{{'1', '1', '1', '1', '0'},
	//	{'1', '1', '0', '1', '0'},
	//	{'1', '1', '0', '0', '0'},
	//	{'0', '0', '0', '0', '0'}}))
	//fmt.Println(numIslands([][]byte{{'1', '1', '0', '0', '0'},
	//	{'1', '1', '0', '0', '0'},
	//	{'0', '0', '1', '0', '0'},
	//	{'0', '0', '0', '1', '1'}}))
	fmt.Println(numIslands0([][]byte{{'1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '0', '1', '1'}, {'0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '0'}, {'1', '0', '1', '1', '1', '0', '0', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, {'1', '0', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '0', '1', '1', '1', '0', '1', '1', '1'}, {'0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '0', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '0', '1', '1'}, {'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, {'0', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1'}, {'1', '0', '1', '1', '1', '1', '1', '0', '1', '1', '1', '0', '1', '1', '1', '1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '0'}, {'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '0', '0'}, {'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}}))
}

//并查集
func numIslands(grid [][]byte) int {
	optionX := []int{0, 1, 0, -1}
	optionY := []int{1, 0, -1, 0}
	union := ConstructorUnionFind200(grid)
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				for k := 0; k < 4; k++ {
					new := []int{i + optionX[k], j + optionY[k]}
					if 0 <= new[0] && new[0] < m && 0 <= new[1] && new[1] < n && grid[new[0]][new[1]] == '1' {
						if !union.Same(i*n+j, new[0]*n+new[1]) {
							union.Union(i*n+j, new[0]*n+new[1])
						}
					}
				}
			}
		}
	}
	return union.count
}

type UnionFind200 struct {
	count  int
	parent []int
	rank   []int
	//注意
	//rank数组主要是为了在union的时候，通过查询rank数组，将rank低的合并到rank高的树上，防止多次union后树高度太高find效率降低。
	//此方式和路径压缩按情况,取其一.
}

func ConstructorUnionFind200(grid [][]byte) UnionFind200 {
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
	return UnionFind200{
		count:  count,
		parent: parent,
		rank:   rank,
	}
}

func (u *UnionFind200) Find(index int) int {
	if u.parent[index] == -1 {
		return -1
	}
	root := index
	for root != u.parent[root] {
		root = u.parent[root]
	}
	//注意
	//路径压缩,方便快速找到root节点
	//此方式和rank按情况,取其一.
	i, j := index, 0
	for root != u.parent[i] {
		j = u.parent[i]
		u.parent[i] = root
		i = j
	}
	return root
}

func (u *UnionFind200) Union(x, y int) {
	xp := u.Find(x)
	yp := u.Find(y)
	if xp == yp {
		return
	}
	if u.rank[x] < u.rank[y] {
		u.parent[xp] = yp
	} else {
		u.parent[yp] = xp
	}
	if u.rank[x] == u.rank[y] {
		u.rank[x]++
	}
	u.count--
}

func (u *UnionFind200) Same(x, y int) bool {
	xp := u.Find(x)
	yp := u.Find(y)
	return xp == yp
}

//BFSDFS
func numIslands0(grid [][]byte) int {
	optionX := []int{0, 1, 0, -1}
	optionY := []int{1, 0, -1, 0}
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	count := 0
	queue := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				queue = append(queue, []int{i, j})
				grid[i][j] = '0'
				for len(queue) > 0 {
					now := queue[0]
					queue = queue[1:]
					for k := 0; k < 4; k++ {
						new := []int{now[0] + optionX[k], now[1] + optionY[k]}
						if 0 <= new[0] && new[0] < m && 0 <= new[1] && new[1] < n && grid[new[0]][new[1]] == '1' {
							queue = append(queue, new)
							grid[new[0]][new[1]] = '0'
						}
					}
				}
			}
		}
	}
	return count
}
