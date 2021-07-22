package Backpack

import (
	"testing"
)

/*
给出一个都是正整数的数组 nums，其中没有重复的数。从中找出所有的和为
target 的组合个数。
样例
给出 nums = [1, 2, 4], target = 4
可能的所有组合有：
[1, 1, 1, 1]
[1, 1, 2]
[1, 2, 1]
[2, 1, 1]
[2, 2]
[4]
返回 6
*/

//完全背包
//564 · 组合总和 IV
//https://www.lintcode.com/problem/564/
func Test_backPackVI(t *testing.T) {
	println(backPackVI([]int{1, 2, 4}, 4))
}

func backPackVI(nums []int, target int) int {
	//容量为f[i]时的方案数
	f := make([]int, target+1)
	f[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				f[i] += f[i-num]
			}
		}
	}
	return f[target]
}
func backPackVI1(nums []int, target int) int {
	if target == 0 {
		return 1
	}
	if target < 0 {
		return 0
	}
	ans := 0
	for _, num := range nums {
		ans += backPackVI(nums, target-num)
	}
	return ans
}

//func backPackVI00(nums []int, target int) int {
//	n := len(nums)
//	//f[i][j]表示前i个物品,使用i的情况下,背包容量为j,的方案数
//	f := make([][]int, n+1)
//	for i := range f {
//		f[i] = make([]int, target+1)
//		f[i][0] = 1
//	}
//	for i := 0; i < n; i++ {
//		for j := 0; j <= target; j++ {
//			if j-nums[i] >= 0 {
//				f[i+1][j] = f[i][j] + f[i][j-nums[i]]
//			} else {
//				f[i+1][j] = f[i][j]
//			}
//		}
//	}
//	ans := 0
//	for i := range f {
//		ans += f[i][target]
//	}
//	return ans
//}

func backPackVI0(nums []int, target int) int {
	n := len(nums)
	//cache := map[string]bool{}
	plan := make([][]int, 0)
	//toString := func(list []int) string {
	//	ans := ""
	//	for _, item := range list {
	//		//ans += fmt.Sprintf("%v_", item)
	//		ans += strconv.Itoa(item) + "_"
	//	}
	//	return ans
	//}
	var dfs func(sum int, list []int)
	dfs = func(sum int, list []int) {
		if sum == target {
			//cache[toString(list)] = true
			plan = append(plan, append(make([]int, 0), list...))
			return
		}
		if sum > target {
			return
		}
		for i := 0; i < n; i++ {
			list = append(list, nums[i])
			sum += nums[i]
			dfs(sum, list)
			list = list[:len(list)-1]
			sum -= nums[i]
		}
	}
	dfs(0, make([]int, 0))
	return len(plan)
}
