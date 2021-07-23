package Backpack

import (
	"testing"
)

/*
描述
给一些不同价值和数量的硬币。找出[1，n]范围内的总值有多少种形成方式？

样例 1:
	输入:
		n = 5
		value = [1,4]
		amount = [2,1]
	输出:  4

	解释:
	可以组合出1，2，4，5

样例 2:
	输入:
		n = 10
		value = [1,2,4]
		amount = [2,1,1]
	输出:  8

	解释:
	可以组合出1-8所有数字。
*/

//

func Test_backPackVIII(t *testing.T) {
	println(backPackVIII(5, []int{1, 4}, []int{2, 1}))
}

func backPackVIII(n int, value []int, amount []int) int {
	//f[i]表示,i被表示出方案数
	f := make([]int, n+1)
	m := len(value)
	f[0] = 1
	for i := 0; i < m; i++ {
		cost := value[i] * amount[i]
		if cost >= n {
			//完全背包
			for j := value[i]; j <= n; j++ {
				f[j] += f[j-value[i]]
			}
		} else {
			//非完全
			count := amount[i]
			for j := 1; j < count; j <<= 1 {
				for k := n; k >= j*value[i]; k-- {
					f[k] += f[k-j*value[i]]
				}
				count -= j
			}
			if count != 0 {
				for k := n; k >= count*value[i]; k-- {
					f[k] += f[k-count*value[i]]
				}
			}
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		if f[i] != 0 {
			ans++
		}
	}
	return ans
}

func backPackVIII0(n int, value []int, amount []int) int {
	//f[i]表示,i被表示出方案数
	f := make([]int, n+1)
	m := len(value)
	f[0] = 1
	for i := 0; i < m; i++ {
		for j := n; j >= value[i]; j-- {
			for k := 1; k <= amount[i] && j-k*value[i] >= 0; k++ {
				f[j] += f[j-k*value[i]]
			}
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		if f[i] != 0 {
			ans++
		}
	}
	return ans
}

//先数量后价格不行,错误答案
func backPackVIII0XX(n int, value []int, amount []int) int {
	//f[i]表示,i被表示出方案数
	f := make([]int, n+1)
	m := len(value)
	f[0] = 1
	for i := 0; i < m; i++ {
		for j := 1; j <= amount[i]; j++ {
			for k := n; k >= value[i] && k-j*value[i] >= 0; k-- {
				f[k] += f[k-j*value[i]]
			}
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		if f[i] != 0 {
			ans++
		}
	}
	return ans
}
