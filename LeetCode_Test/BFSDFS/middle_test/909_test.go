package middle_test

import "testing"

func Test_snakesAndLadders(t *testing.T) {
	board := [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1}}
	fmt.Println(snakesAndLadders(board))
}

func snakesAndLadders(board [][]int) int {
	n := len(board)
	isLeft := 1
	list := make([]int, 0, n*n+1)
	list = append(list, -1)
	for i := n - 1; i >= 0; i-- {
		if isLeft == 1 {
			for j := 0; j < n; j++ {
				list = append(list, board[i][j])
			}
		} else {
			for j := n - 1; j >= 0; j-- {
				list = append(list, board[i][j])
			}
		}
		isLeft = 1 ^ isLeft
	}
	queue := make([]int, 0)
	queue = append(queue, 1)
	visited := map[int]int{}
	visited[1] = 0
	for len(queue) != 0 {
		now := queue[0]
		queue = queue[1:]
		for i := 1; i <= 6; i++ {
			next := now + i
			if next > n*n {
				break
			}
			if list[next] > 0 {
				next = list[next]
			}
			if next == n*n {
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
