package Tree

import (
	"fmt"
	"testing"
)

//所有可能的路径
//https://leetcode.cn/problems/all-paths-from-source-to-target/description/
func allPathsSourceTarget(graph [][]int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	dfs(graph, &ans, path, 0)
	return ans
}

func dfs(graph [][]int, ans *[][]int, path []int, point int) {
	path = append(path, point)

	for _, i := range graph[point] {
		dfs(graph, ans, path, i)
	}
	if point == len(graph)-1 {
		*ans = append(*ans, append([]int{}, path...))
	}
	path = path[:len(path)-1]
}

func Test_0(t *testing.T) {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
}
