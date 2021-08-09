package middle_test_test

import "testing"

func Test_eventualSafeNodes(t *testing.T) {
	println(eventualSafeNodes([][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}))
}

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	cache := map[int]int{}
	ans := make([]int, 0)
	//初始化
	for i := 0; i < n; i++ {
		cache[i] = 0
		if safe(graph, cache, i) {
			ans = append(ans, i)
		}
	}
	return ans
}

func safe(graph [][]int, cache map[int]int, num int) bool {
	if cache[num] != 0 {
		return cache[num] == 2
	}
	cache[num] = 1
	for _, i := range graph[num] {
		if !safe(graph, cache, i) {
			return false
		}
	}
	cache[num] = 2
	return true
}
