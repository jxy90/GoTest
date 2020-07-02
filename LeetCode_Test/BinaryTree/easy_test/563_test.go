package easy_test

import (
	"math"
	"testing"
)

var tilt int

func findTilt(root *TreeNode) int {
	tilt = 0
	findTiltDFS(root)
	return tilt
}

func findTiltDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	Left := findTiltDFS(root.Left)
	Right := findTiltDFS(root.Right)
	tilt += int(math.Abs(float64(Left) - float64(Right)))
	return Left + Right + root.Val
}

//func findTiltDFSVisit(root *TreeNode) {
//	if root.Left == nil && root.Right == nil {
//		tilt += 0
//	} else if root.Left == nil {
//		tilt += root.Right.Val
//	} else if root.Right == nil {
//		tilt += root.Left.Val
//	} else {
//		tilt += int(math.Abs(float64(root.Left.Val) - float64(root.Right.Val)))
//	}
//}

func Test_findTilt(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		//Right: &TreeNode{
		//	Val: 3,
		//},
	}
	data := findTilt(root)
	println(data)
}
