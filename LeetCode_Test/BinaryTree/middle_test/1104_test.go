package middle_test

import (
	"fmt"
	"math"
	"testing"
)

func Test_pathInZigZagTree(t *testing.T) {
	fmt.Println(CalNum(0)) //1
	fmt.Println(CalNum(1)) //5
	fmt.Println(CalNum(2)) //11
	fmt.Println(CalNum(3)) //23
	fmt.Println(CalNum(4))
	fmt.Println(int(math.Log2(1)))
}

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
	LeftNum := int(math.Pow(2, float64(level)))
	RightNum := LeftNum + int(math.Pow(2, float64(level))) - 1
	return LeftNum + RightNum
}
