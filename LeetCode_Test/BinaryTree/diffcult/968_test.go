package diffcult_test

import (
	"math"
	"testing"
)

func minCameraCover(root *TreeNode) int {
	set = map[*TreeNode]int{}
	//把nil加入set中，这样末尾检点就不会设置监视器
	set[nil] = 1
	ans = 0
	minCameraCoverDFS(root, nil)
	return ans
}


var set map[*TreeNode]int
var ans int

func minCameraCoverDFS(root, parent *TreeNode) {
	if root != nil {
		minCameraCoverDFS(root.Left, root)
		minCameraCoverDFS(root.Right, root)

		//贪心法 后序遍历
		//1。当没有父节点并且自己没有被监控，把自己加入，并把nil，左，右 节点加入set
		//2。当自己的左子节点或者右子节点没有被监控的时，给自己添加监视器，并把父，左，右 节点加入set
		//因为set时map 重复设置没有关系。
		if parent == nil && set[root] == 0 || set[root.Right] == 0 || set[root.Left] == 0 {
			ans++
			set[parent] = 1
			set[root] = 1
			set[root.Left] = 1
			set[root.Right] = 1
		}
	}
}

func minCameraCover2(root *TreeNode) int {
	ans := minCameraCoverSolve(root)
	return int(math.Min(ans[1],ans[2]))
}

func minCameraCoverSolve(root *TreeNode) []float64 {
	if root == nil {
		return []float64{0, 0, 9999}
	}
	left := minCameraCoverSolve(root.Left)
	right := minCameraCoverSolve(root.Right)

	minLeft12 := math.Min(left[1], left[2])
	minRight12 := math.Min(right[1], right[2])

	d0 := left[1] + right[1]
	d1 := math.Min(left[2]+minRight12, right[2]+minLeft12)
	d2 := 1 + math.Min(left[0], minLeft12) + math.Min(right[0], minRight12)
	return []float64{d0, d1, d2}
}

func Test_minCameraCover(t *testing.T) {
	set = map[*TreeNode]int{}
	root := &TreeNode{
		Val: 0,
	}
	if set[root] == 0 {
		println(0)
	}
	println(2)

	if set[nil] == 0 {
		println(0)
	}
	set[nil] = 1
	if set[nil] == 1 {
		println(1)
	}

}
