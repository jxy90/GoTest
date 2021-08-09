package middle_test_test

import "testing"

func Test_shortestPathLength(t *testing.T) {
	println(shortestPathLength([][]int{{1, 2, 3}, {0}, {0}, {0}}))
}

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	type Node struct {
		node, visited, length int
	}
	//history[i][j]记录第i个数,状态j是否被访问过
	history := make([][]bool, n)
	q := make([]Node, 0)
	for i := range history {
		history[i] = make([]bool, 1<<n)
		//设置自己的访问状态
		history[i][1<<i] = true
		q = append(q, Node{
			node:    i,
			visited: 1 << i,
			length:  0,
		})
	}

	for {
		now := q[0]
		q = q[1:]
		if now.visited == (1<<n)-1 {
			return now.length
		}
		for _, next := range graph[now.node] {
			visitedVal := now.visited | (1 << next)
			//next数的visitedVal状态被访问过,就跳过
			if !history[next][visitedVal] {
				q = append(q, Node{
					node:    next,
					visited: visitedVal,
					length:  now.length + 1,
				})
				history[next][visitedVal] = true
			}
		}
	}

}
