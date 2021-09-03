package middle_test

import (
	"fmt"
	"math"
	"testing"
)

func Test_pathInZigZagTree0(t *testing.T) {
	//fmt.Println(int(math.Log2(1)))
	//fmt.Println(int(math.Log2(3)))
	//fmt.Println(int(math.Log2(4)))
	//fmt.Println(int(math.Log2(14)))
	fmt.Println(pathInZigZagTree_0(14))
}

func pathInZigZagTree_0(label int) []int {
	var getLevelMinLabel func(int) int
	getLevelMinLabel = func(i int) int {
		return int(math.Pow(2, float64(i)))
	}
	//获取当前层数
	level := int(math.Log2(float64(label)))
	ans := make([]int, level+1)
	ans[level] = label
	level--
	for level >= 0 {
		//当前行的最大和最小值
		min, max := getLevelMinLabel(level), getLevelMinLabel(level+1)-1
		sum := min + max
		//当都是正序排列的时候,父节点的值=子节点/2
		label = label / 2
		//由于奇数行,偶数行,顺序不同,父节点的值=正序排列的值的对应值,即为下面
		label = sum - label
		//加入数组
		ans[level] = label
		//下一层
		level--
	}
	return ans
}
