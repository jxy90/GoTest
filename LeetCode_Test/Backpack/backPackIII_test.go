package Backpack

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

/*
给定 n 种物品, 每种物品都有无限个. 第 i 个物品的体积为 A[i], 价值为 V[i].

再给定一个容量为 m 的背包. 问可以装入背包的最大价值是多少?
*/

//完全背包
func Test_backPackIII(t *testing.T) {
	fmt.Println(backPackIII(10, []int{2, 3, 5, 7}, []int{1, 5, 2, 4}))
	fmt.Println(backPackIII(100, []int{}, []int{}))
}

func backPackIII0(m int, A, V []int) int {
	n := len(A)
	//f[i]表示背包容量i时,最多能装的值
	f := make([]int, m+1)
	for i := 0; i < n; i++ {
		//完全背包,正序遍历,物品可复用
		for j := A[i]; j <= m; j++ {
			f[j] = CommonUtil.Max(f[j], f[j-A[i]]+V[i])
		}
	}
	return f[m]
}

func backPackIII(m int, A, V []int) int {
	n := len(A)
	//f[i][j]表示前i个物品,背包容量为j,最多能装的值
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= m; j++ {
			//count表示第i件物品能用的最大数量
			count := m / A[i]
			//循环count,代表每次放入k*A[i]大小的物品
			for k := 0; k <= count; k++ {
				if j-k*A[i] >= 0 {
					f[i+1][j] = CommonUtil.Max(f[i][j], f[i][j-k*A[i]]+k*V[i])
				} else {
					f[i+1][j] = f[i][j]
				}
			}
		}
	}
	return f[n][m]
}
