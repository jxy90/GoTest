package BackPack

import (
	"fmt"
	"testing"
)

//198. 打家劫舍
func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[n]
}

//213. 打家劫舍 II
func rob2(nums []int) int {
	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

//337. 打家劫舍 III
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob3(root *TreeNode) int {
	memoCan := make(map[*TreeNode]int)
	memoCant := make(map[*TreeNode]int)
	var helper func(root *TreeNode, can bool) int
	helper = func(root *TreeNode, can bool) int {
		if root == nil {
			return 0
		}
		if !can {
			if val, ok := memoCant[root]; ok {
				return val
			}
			res := helper(root.Left, !can) + helper(root.Right, !can)
			memoCant[root] = res
			return res
		}
		if val, ok := memoCan[root]; ok {
			return val
		}
		res := max(root.Val+rob3Helper(root.Left, !can)+rob3Helper(root.Right, !can), rob3Helper(root.Left, can)+rob3Helper(root.Right, can))
		memoCan[root] = res
		return res
	}
	return max(helper(root, true), helper(root, false))
}
func rob3Helper(root *TreeNode, can bool) int {
	if root == nil {
		return 0
	}
	if !can {
		return rob3Helper(root.Left, !can) + rob3Helper(root.Right, !can)
	}
	return max(root.Val+rob3Helper(root.Left, !can)+rob3Helper(root.Right, !can), rob3Helper(root.Left, can)+rob3Helper(root.Right, can))
}

func Test_9(t *testing.T) {
	//fmt.Println(rob([]int{1, 2, 3, 1}))
	//fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob2([]int{2, 3, 2}))
	fmt.Println(rob2([]int{1, 2, 3, 1}))
	fmt.Println(rob2([]int{1, 2, 3}))
}
