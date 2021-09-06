package DP_test

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_profitableSchemes(t *testing.T) {
	//100
	//100
	//[]
	//[]
	fmt.Println(profitableSchemes(100, 100, []int{2, 5, 36, 2, 5, 5, 14, 1, 12, 1, 14, 15, 1, 1, 27, 13, 6, 59, 6, 1, 7, 1, 2, 7, 6, 1, 6, 1, 3, 1, 2, 11, 3, 39, 21, 20, 1, 27, 26, 22, 11, 17, 3, 2, 4, 5, 6, 18, 4, 14, 1, 1, 1, 3, 12, 9, 7, 3, 16, 5, 1, 19, 4, 8, 6, 3, 2, 7, 3, 5, 12, 6, 15, 2, 11, 12, 12, 21, 5, 1, 13, 2, 29, 38, 10, 17, 1, 14, 1, 62, 7, 1, 14, 6, 4, 16, 6, 4, 32, 48}, []int{21, 4, 9, 12, 5, 8, 8, 5, 14, 18, 43, 24, 3, 0, 20, 9, 0, 24, 4, 0, 0, 7, 3, 13, 6, 5, 19, 6, 3, 14, 9, 5, 5, 6, 4, 7, 20, 2, 13, 0, 1, 19, 4, 0, 11, 9, 6, 15, 15, 7, 1, 25, 17, 4, 4, 3, 43, 46, 82, 15, 12, 4, 1, 8, 24, 3, 15, 3, 6, 3, 0, 8, 10, 8, 10, 1, 21, 13, 10, 28, 11, 27, 17, 1, 13, 10, 11, 4, 36, 26, 4, 2, 2, 2, 10, 0, 11, 5, 22, 6}))
	fmt.Println(profitableSchemes(5, 3, []int{2, 2}, []int{2, 3}))
	fmt.Println(profitableSchemes(10, 5, []int{2, 2}, []int{2, 3}))
}

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	mod := 1000000000 + 7
	m := len(group)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, minProfit+1)
		f[i][0] = 1
	}
	//f[i][j]前i个组获得j的利润的方法数量
	for i := 1; i <= m; i++ {
		a, b := group[i-1], profit[i-1]
		for j := n; j >= a; j-- {
			for k := minProfit; k >= 0; k-- {
				f[j][k] += f[j-a][CommonUtil.Max(k-b, 0)]
				if f[j][k] >= mod {
					f[j][k] -= mod
				}
			}
		}
	}

	return f[n][minProfit]
}
