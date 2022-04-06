package middle_test

import (
	"testing"
)

func Test_310(t *testing.T) {
	t.Log(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
}

func findMinHeightTrees(n int, edges [][]int) (ans []int) {
	if n == 1 {
		return []int{0}
	}
	graph := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	parents := make([]int, n)
	dfs := func(start int) (x int) {
		visited := map[int]bool{}
		visited[start] = true
		q := []int{start}
		for len(q) > 0 {
			x, q = q[0], q[1:]
			for _, y := range graph[x] {
				if !visited[y] {
					visited[y] = true
					parents[y] = x
					q = append(q, y)
				}
			}
		}
		return
	}
	//x和y是相距最远的两个点
	x := dfs(0)
	y := dfs(x)
	//记录y到x经历的点,路径
	path := make([]int, 0)
	parents[x] = -1
	for y != -1 {
		path = append(path, y)
		y = parents[y]
	}
	//答案就是这个路径最中间的一个或者两个点
	m := len(path)
	if m%2 == 0 {
		return []int{path[m/2-1], path[m/2]}
	}
	return []int{path[m/2]}
}
