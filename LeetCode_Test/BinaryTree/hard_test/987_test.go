package hard_test

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func Test_verticalTraversal(t *testing.T) {
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(verticalTraversal(root))
}

type data struct{ col, row, val int }

func verticalTraversal0(root *TreeNode) [][]int {
	nodes := []data{}
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, row, col int) {
		if node == nil {
			return
		}
		nodes = append(nodes, data{col, row, node.Val})
		dfs(node.Left, row+1, col-1)
		dfs(node.Right, row+1, col+1)
	}
	dfs(root, 0, 0)

	sort.Slice(nodes, func(i, j int) bool {
		a, b := nodes[i], nodes[j]
		return a.col < b.col || a.col == b.col && (a.row < b.row || a.row == b.row && a.val < b.val)
	})

	lastCol := math.MinInt32
	ans := make([][]int, 0)
	for _, node := range nodes {
		if node.col != lastCol {
			lastCol = node.col
			ans = append(ans, nil)
		}
		ans[len(ans)-1] = append(ans[len(ans)-1], node.val)
	}
	return ans
}

func verticalTraversal(root *TreeNode) [][]int {
	var nodes []data
	//递归遍历
	var loop func(*TreeNode, int, int)
	loop = func(node *TreeNode, left, right int) {
		if node == nil {
			return
		}
		nodes = append(nodes, data{
			col: right,
			row: left,
			val: node.Val,
		})
		loop(node.Left, left+1, right-1)
		loop(node.Right, left+1, right+1)
	}
	loop(root, 0, 0)

	sort.Slice(nodes, func(i, j int) bool {
		a, b := nodes[i], nodes[j]
		return a.col < b.col || a.col == b.col && (a.row < b.row || a.row == b.row && a.val < b.val)
	})
	lastCol := math.MinInt32
	ans := make([][]int, 0)
	for _, node := range nodes {
		if node.col != lastCol {
			lastCol = node.col
			ans = append(ans, nil)
		}
		ans[len(ans)-1] = append(ans[len(ans)-1], node.val)
	}
	return ans
}

func verticalTraversal0X(root *TreeNode) [][]int {
	//cache := map[int][]int{}
	cache := map[int]map[int][]int{}
	//递归遍历
	var loop func(root *TreeNode, left, right int)
	loop = func(root *TreeNode, left, right int) {
		if root == nil {
			return
		}
		cache[left][right] = append(cache[left][right], root.Val)
		loop(root.Left, left+1, right-1)
		loop(root.Right, left+1, right+1)
	}
	//加入cache
	loop(root, 0, 0)
	//列值排序
	temp := make([]int, 0)
	for _, rowV := range cache {
		//cache内部数组排序
		for columnK := range rowV {
			sort.Ints(rowV[columnK])
			temp = append(temp, columnK)
		}
	}
	sort.Ints(temp)
	//按顺序赋值
	ans := make([][]int, len(temp))
	//for i := range temp {
	//ans[i] = cache[temp[i]]
	//}
	return ans
}
