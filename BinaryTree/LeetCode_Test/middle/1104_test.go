package middle_test

import (
	"math"
	"testing"
)

func pathInZigZagTree(label int) []int {
	var ls []int
	level := int(math.Log2(float64(label)))
	for ; level > 0; level-- {
		ls = append([]int{label}, ls...)
		if level%2 == 0 {
			label = CalNum(level-1) - label/2
		} else {
			label = (CalNum(level) - label) / 2
		}
	}
	if level == 0 {
		ls = append([]int{1}, ls...)
	}
	return ls
}

//计算第level层的左右节点的和，方便求出它的对称节点
func CalNum(level int) int {
	if level == 0 {
		return 1
	}
	leftNum := int(math.Pow(2, float64(level)))
	rightNum := leftNum + int(math.Pow(2, float64(level))) - 1
	return leftNum + rightNum
}

func Test_pathInZigZagTree(t *testing.T) {
	println(CalNum(0)) //1
	println(CalNum(1)) //5
	println(CalNum(2)) //11
	println(CalNum(3)) //23
	println(CalNum(4))
	println(int(math.Log2(1)))
}
