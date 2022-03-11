package middle_test

import (
	"fmt"
	"testing"
)

func Test_1339(t *testing.T) {
	//fmt.Println(maxProduct())
}

func maxProduct(root *TreeNode) int {
	var res int
	var sum int
	var getSum func(root *TreeNode) int
	getSum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftSum := getSum(root.Left)
		rightSum := getSum(root.Right)
		rootSum := leftSum + rightSum + root.Val
		res = max(res, rootSum*(sum-rootSum))
		return rootSum
	}
	sum = getSum(root)
	getSum(root)
	fmt.Println(sum)
	return res % (1e9 + 7)
}
