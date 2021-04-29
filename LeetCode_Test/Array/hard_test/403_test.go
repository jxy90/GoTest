package hard_test

import "testing"

func Test_canCross(t *testing.T) {
	println(canCross2([]int{0, 1, 3, 5, 6, 8, 12, 17}))
	println(canCross2([]int{0, 1, 2, 3, 4, 8, 9, 11}))
}

//dp[i][j]代表：第i个为止，通过j到达
//func canCross2(stones []int) bool {
//	if len(stones) < 2 || stones[1]-stones[0] != 1 {
//		return false
//	}
//
//	has := map[int]bool{}
//	for i := range stones {
//		has[stones[i]] = true
//	}
//
//	n := len(stones)
//	dp := make([][]int, n+1)
//	for i := range dp {
//		dp[i] = make([]int, n+1)
//	}
//	dp[1][1] = 1
//
//}
//BFS
func canCross2(stones []int) bool {
	if len(stones) < 2 || stones[1]-stones[0] != 1 {
		return false
	}
	n := len(stones)
	choice := []int{-1, 0, 1}

	has := map[int]bool{}
	for i := range stones {
		has[stones[i]] = true
	}

	visited := make([][]bool, n+1)
	for i := range visited {
		visited[i] = make([]bool, n+1)
	}
	visited[1][1] = true
	queue := make([][]int, 0)
	queue = append(queue, []int{1, 1})
	for len(queue) != 0 {
		position := queue[0][0]
		step := queue[0][1]
		if position == stones[n-1] {
			return true
		}
		queue = queue[1:]
		for _, value := range choice {
			newStep := step + value
			newPosition := position + newStep
			if newStep > 0 && has[newPosition] && visited[newPosition][newStep] == false {
				visited[newPosition][newStep] = true
				queue = append(queue, []int{newPosition, newStep})
			}
		}
	}
	return false
}

//dfs
func canCross(stones []int) bool {
	if len(stones) < 2 || stones[1]-stones[0] != 1 {
		return false
	}
	choice := []int{-1, 0, 1}

	has := map[int]bool{}
	for i := range stones {
		has[stones[i]] = true
	}

	ans := false
	var canCrossDFS func(stones []int, val, k int)
	canCrossDFS = func(stones []int, val, k int) {

		if val == stones[len(stones)-1] {
			ans = true
			return
		}
		for i := range choice {
			x := k + choice[i]
			if x > 0 && has[val+x] {
				canCrossDFS(stones, val+x, x)
			}
		}
	}
	canCrossDFS(stones, stones[1], stones[1]-stones[0])
	return ans
}
