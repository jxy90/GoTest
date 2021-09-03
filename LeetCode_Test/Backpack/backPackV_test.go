package Backpack

import (
	"testing"
)

/*
给出 n 个物品, 以及一个数组, nums[i] 代表第i个物品的大小, 保证大小均为正数, 正整数 target 表示背包的大小, 找到能填满背包的方案数。
每一个物品只能使用一次
*/

//01背包
func Test_backPackV(t *testing.T) {
	fmt.Println(backPackV([]int{1, 2, 3, 3, 7}, 7))
}

func backPackV(nums []int, target int) int {
	n := len(nums)
	//f[i]表示背包容量i时,装满的方案数
	f := make([]int, target+1)
	f[0] = 1
	for i := 0; i < n; i++ {
		for j := target; j >= nums[i]; j-- {
			f[j] += f[j-nums[i]]
		}
	}
	return f[target]
}

func backPackV0(nums []int, target int) int {
	n := len(nums)
	//f[i][j]表示前i个物品,背包容量为j,装满的方案数
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, target+1)
		f[i][0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= target; j++ {
			if j-nums[i] >= 0 {
				f[i+1][j] = f[i][j] + f[i][j-nums[i]]
			} else {
				f[i+1][j] = f[i][j]
			}
		}
	}
	return f[n][target]
}
