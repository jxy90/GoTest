package middle_test

import (
	"fmt"
	"testing"
)

func Test_2049(t *testing.T) {
	fmt.Println(countHighestScoreNodes([]int{-1, 2, 0, 2, 0}))
}

//邻接表
var tree [][]int
var scores map[int][]int

func countHighestScoreNodes(parents []int) int {
	scores = map[int][]int{}
	buildTree(parents)
	countNode(0)
	maxScore := 0
	temp := []int{}
	for k, ints := range scores {
		if k > maxScore {
			maxScore = k
			temp = ints
		}
	}
	fmt.Println(temp)
	return len(temp)
}

func countNode(root int) int {
	if root == -1 {
		return 0
	}
	n := len(tree)
	leftCount := countNode(tree[root][0])
	rightCount := countNode(tree[root][1])

	otherCount := n - leftCount - rightCount - 1
	score := max(leftCount, 1) * max(rightCount, 1) * max(otherCount, 1)
	scores[score] = append(scores[score], root)
	return leftCount + rightCount + 1
}

func max(args ...int) int {
	max := args[0]
	for _, arg := range args {
		if arg > max {
			max = arg
		}
	}
	return max
}

func buildTree(parent []int) {
	n := len(parent)
	tree = make([][]int, n)
	for i := range tree {
		tree[i] = make([]int, 2)
		for j := range tree[i] {
			tree[i][j] = -1
		}
	}
	for i := 1; i < n; i++ {
		p := parent[i]
		if tree[p][0] == -1 {
			tree[p][0] = i
		} else {
			tree[p][1] = i
		}
	}
}
