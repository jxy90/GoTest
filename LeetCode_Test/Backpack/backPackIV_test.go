package Backpack

import (
	"fmt"
	"testing"
)

/*
给出 n 个物品, 以及一个数组, nums[i]代表第i个物品的大小, 保证大小均为正数并且没有重复, 正整数 target 表示背包的大小, 找到能填满背包的方案数。
每一个物品可以使用无数次
*/

//完全背包
func Test_backPackIV(t *testing.T) {
	fmt.Println(backPackIV([]int{2, 3, 6, 7}, 7))
	fmt.Println(backPackIV([]int{2, 3, 4, 5}, 7))
}

func backPackIV(nums []int, target int) int {
	n := len(nums)
	//f[i]表示背包容量i时,装满的方案数
	f := make([]int, target+1)
	f[0] = 1
	for i := 0; i < n; i++ {
		//完全背包,正序遍历,物品可复用
		for j := nums[i]; j <= target; j++ {
			f[j] = f[j] + f[j-nums[i]]
		}
	}
	return f[target]
}

func backPackIV0(nums []int, target int) int {
	n := len(nums)
	//f[i][j]表示前i个物品,背包容量为j,装满方案数
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, target+1)
		f[i][0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 1; j <= target; j++ {
			//count表示第i件物品能用的最大数量
			count := target / nums[i]
			//循环count,代表每次放入k*A[i]大小的物品
			for k := 0; k <= count; k++ {
				if k == 0 {
					f[i+1][j] = f[i][j]
				} else if j-k*nums[i] >= 0 {
					f[i+1][j] += f[i][j-k*nums[i]]
				}
			}
		}
	}
	return f[n][target]
}
