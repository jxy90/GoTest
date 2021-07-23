package Backpack

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

/*
描述
你总共有n元，商人总共有三种商品，它们的价格分别是150元,250元,350元，三种商品的数量可以认为是无限多的，购买完商品以后需要将剩下的钱给商人作为小费，求最少需要给商人多少小费

样例
样例 1:

输入:  n = 900
输出:  0
样例 2:

输入:  n = 800
输出:  0
*/

//

func Test_backpackX(t *testing.T) {
	println(backPackX(900))
	println(backPackX(800))
	println(backPackX(1000))
}

func backPackX(n int) int {
	f := make([]int, n+1)
	for i := range f {
		f[i] = n
	}
	ints := []int{150, 250, 350}
	m := len(ints)
	for i := 0; i < m; i++ {
		for j := ints[i]; j <= n; j++ {
			f[j] = CommonUtil.Min(f[j], f[j-ints[i]]-ints[i])
		}
	}
	return f[n]
}
