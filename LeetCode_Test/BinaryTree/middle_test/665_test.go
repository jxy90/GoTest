package middle_test

import (
	"math"
	"strconv"
)

func printTree(root *TreeNode) [][]string {
	//深度和行数
	m := printTreeDepth(root)
	//列数
	n := math.Pow(2, m) - 1
	//二维数组Slice初始化
	var result = make([][]string, 0)
	for i := 0; i < int(m); i++ {
		result = append(result, make([]string, int(n)))
	}
	//二分法添加
	printTreeFill(root, 0, int(n)-1, 0, &result)
	return result
}

func printTreeFill(root *TreeNode, L, R, depth int, result *[][]string) {
	if root == nil {
		return
	}
	(*result)[depth][(L+R)/2] = strconv.Itoa(root.Val)
	printTreeFill(root.Left, L, (L+R)/2-1, depth+1, result)
	printTreeFill(root.Right, (R+L)/2+1, R, depth+1, result)
}

func printTreeDepth(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	return math.Max(printTreeDepth(root.Left), printTreeDepth(root.Right)) + 1
}
