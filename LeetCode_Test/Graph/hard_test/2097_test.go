package middle_test_test

import (
	"fmt"
	"testing"
)

func Test_validArrangement(t *testing.T) {
	//fmt.Println(validArrangement([][]int{{5, 1}, {4, 5}, {11, 9}, {9, 4}}))
	//fmt.Println(validArrangement([][]int{{1, 3}, {2, 1}, {3, 2}}))
	//fmt.Println(validArrangement([][]int{{1, 2}, {2, 1}, {1, 3}}))
	fmt.Println(validArrangement([][]int{{8, 5}, {8, 7}, {0, 8}, {0, 5}, {7, 0}, {5, 0}, {0, 7}, {8, 0}, {7, 8}}))
	//
}
func validArrangement(pairs [][]int) [][]int {
	n := len(pairs)
	//定义图和度
	graph := map[int][]int{}
	degIN, degOut, deg := map[int]int{}, map[int]int{}, map[int]int{}
	for _, pair := range pairs {
		//if degIN[pair[1]] > 0 {
		//	graph[pair[0]] = append([]int{pair[1]}, graph[pair[0]]...)
		//} else {
		graph[pair[0]] = append(graph[pair[0]], pair[1])
		//}
		degOut[pair[0]]--
		degIN[pair[1]]++
		deg[pair[0]]--
		deg[pair[1]]++
	}
	//判断是否存在入库为-1的点,存在将其当做起点
	for k, v := range deg {
		if v == -1 {
			return validArrangementDFS(n, graph, k)
		}
	}
	//不存在随便找一点,作为起点
	return validArrangementDFS(n, graph, pairs[0][0])
}
func validArrangementDFS(n int, graph map[int][]int, start int) (res [][]int) {
	if n == len(res) {
		return res
	}
	for n != len(res) {
		end := graph[start][0]
		res = append(res, []int{start, end})
		graph[start] = graph[start][1:]
		start = end
	}
	return res
}
