package easy_test

import (
	"sort"
	"testing"
)

func Test_findSecondMinimumValue(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(findSecondMinimumValue(root))
}

func findSecondMinimumValue(root *TreeNode) int {
	if root.Left == nil {
		return -1
	}
	//nums只记录最小数和第二小数
	nums := []int{0, 0}
	findSecondMinimumValueHelper(root, nums)
	if nums[1] == 0 {
		return -1
	}
	if nums[0] > nums[1] {
		return nums[0]
	}
	return nums[1]
}

//循环取所有数
func findSecondMinimumValueHelper(root *TreeNode, nums []int) {
	if root == nil {
		return
	}
	findSecondMinimumValueNum(nums, root.Val)
	if root.Left != nil {
		findSecondMinimumValueHelper(root.Left, nums)
		findSecondMinimumValueHelper(root.Right, nums)
	}
}

//更新nums中的较大数
func findSecondMinimumValueNum(nums []int, val int) {
	for i := range nums {
		if nums[i] >= val || nums[i] == 0 {
			nums[i] = val
			return
		}
	}
}

func findSecondMinimumValue0(root *TreeNode) int {
	nums := make(map[int]bool, 0)
	findSecondMinimumValueHelper0(root, nums)
	//快排
	if len(nums) <= 1 {
		return -1
	}
	sortNums := make([]int, 0)
	for key := range nums {
		sortNums = append(sortNums, key)
	}
	sort.Ints(sortNums)
	return sortNums[1]
}

func findSecondMinimumValueHelper0(root *TreeNode, nums map[int]bool) {
	if root == nil {
		return
	}
	nums[root.Val] = true
	if root.Left != nil {
		findSecondMinimumValueHelper0(root.Left, nums)
		findSecondMinimumValueHelper0(root.Right, nums)
	}
}
