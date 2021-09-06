package hard_test

import (
	"fmt"
	"strings"
	"testing"
)

func Test_slidingPuzzle(t *testing.T) {
	fmt.Println(slidingPuzzle([][]int{{1, 2, 3}, {4, 0, 5}}))
}

var neighbors = [6][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}

func slidingPuzzle(board [][]int) int {
	const target = "123450"
	s := make([]byte, 0, 6)
	for _, r := range board {
		for _, v := range r {
			s = append(s, '0'+byte(v))
		}
	}
	start := string(s)
	if start == target {
		return 0
	}
	// 枚举 status 通过一次交换操作得到的状态
	getNext := func(status string) (ret []string) {
		s := []byte(status)
		x := strings.Index(status, "0")
		for _, y := range neighbors[x] {
			s[x], s[y] = s[y], s[x]
			ret = append(ret, string(s))
			s[x], s[y] = s[y], s[x]
		}
		return
	}
	visited := map[string]int{}
	visited[start] = 0
	queue := make([]string, 0)
	queue = append(queue, start)
	for len(queue) != 0 {
		now := queue[0]
		queue = queue[1:]
		nexts := getNext(now)
		for _, next := range nexts {
			if next == target {
				return visited[now] + 1
			}
			if _, ok := visited[next]; !ok {
				visited[next] = visited[now] + 1
				queue = append(queue, next)
			}
		}
	}
	return -1
}
