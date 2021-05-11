package unionFindDisjointSets

import (
	"testing"
)

func Test_findRedundantDirectedConnection(t *testing.T) {
	println(findRedundantDirectedConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
}

func findRedundantDirectedConnection(edges [][]int) []int {
	inDegree := make([]int, 1001)
	for i := range edges {
		inDegree[edges[i][1]]++
	}
	inDegree2Index := make([]int, 0)
	for i := len(edges) - 1; i >= 0; i-- {
		if inDegree[edges[i][1]] == 2 {
			inDegree2Index = append(inDegree2Index, i)
		}
	}
	//入度为2,处理
	if len(inDegree2Index) > 0 {
		for _, index := range inDegree2Index {
			if removeEdgeIsTree(edges, index) {
				return edges[index]
			}
		}
	}
	//有向环处理
	return getRemoveEdge(edges)
}

//判断入度为2的点删除关联能否,成树
func removeEdgeIsTree(edges [][]int, index int) bool {
	father := make([]int, 1001)
	for i := range father {
		father[i] = i
	}
	for i, edge := range edges {
		if i == index {
			continue
		}
		if find685(&father, edge[0]) == find685(&father, edge[1]) {
			return false
		}
		join685(&father, edge[0], edge[1])
	}
	return true
}

//判断删除第一次成环的边
func getRemoveEdge(edges [][]int) []int {
	father := make([]int, 1001)
	for i := range father {
		father[i] = i
	}
	for _, edge := range edges {
		if find685(&father, edge[0]) == find685(&father, edge[1]) {
			return edge
		}
		join685(&father, edge[0], edge[1])
	}
	return nil
}

func find685(f *[]int, x int) int {
	root := x
	for root != (*f)[root] {
		root = (*f)[root]
	}
	//路径压缩
	i, j := x, 0
	for root != (*f)[i] {
		j = (*f)[i]
		(*f)[i] = root
		i = j
	}
	return root
}

func join685(f *[]int, x, y int) {
	if find685(f, x) == find685(f, y) {
		return
	}
	(*f)[find685(f, x)] = find685(f, y)
}
