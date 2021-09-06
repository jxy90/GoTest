package Backpack

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

/*
假设你身上有 n 元，超市里有多种大米可以选择，每种大米都是袋装的，必须整袋购买，给出每种大米的重量，价格以及数量，求最多能买多少公斤的大米

样例 1:

输入:  n = 8, prices = [3,2], weights = [300,160], amounts = [1,6]
输出:  640
解释:  全买价格为2的米。
样例 2:

输入:  n = 8, prices  = [2,4], weight = [100,100], amounts = [4,2 ]
输出:  400
解释:  全买价格为2的米
*/

//

func Test_backPackVII(t *testing.T) {
	fmt.Println(backPackVII(8, []int{3, 2}, []int{300, 160}, []int{1, 6}))
	fmt.Println(backPackVII(8, []int{2, 4}, []int{100, 100}, []int{4, 2}))
	fmt.Println(backPackVII(54, []int{10, 8, 8, 3, 13, 5, 10, 6, 11, 7, 3, 20, 15, 14, 20, 9, 7, 16, 13, 15}, []int{78, 74, 8, 56, 91, 177, 159, 66, 62, 143, 102, 195, 27, 199, 141, 21, 106, 122, 147, 76}, []int{3, 19, 3, 15, 2, 14, 5, 16, 10, 11, 10, 15, 10, 1, 18, 16, 2, 1, 14, 15}))
}

//41ms 4.14MB
//拆分继续优化-01背包部分以减少循环次数
func backPackVII(n int, prices []int, weight []int, amounts []int) int {
	//背包容量为i时,最大能买f[i]
	f := make([]int, n+1)
	m := len(prices)
	for i := 0; i < m; i++ {
		//分情况判断是否完全背包
		cost := prices[i] * amounts[i]
		if cost >= n {
			//n元钱,买不完i商品.完全背包
			for j := prices[i]; j <= n; j++ {
				f[j] = CommonUtil.Max(f[j], f[j-prices[i]]+weight[i])
			}
		} else {
			//01背包
			amount := amounts[i]
			//对数量循环1,2,4,8,16的方式递增,相当于把商品i分成不同的堆,每堆各不相同
			for j := 1; j <= amount; j <<= 1 {
				for k := n; k >= prices[i] && k-j*prices[i] >= 0; k-- {
					f[k] = CommonUtil.Max(f[k], f[k-j*prices[i]]+j*weight[i])
				}
				amount -= j
			}
			if amount != 0 {
				for k := n; k >= prices[i] && k-amount*prices[i] >= 0; k-- {
					//多了个数量循环
					f[k] = CommonUtil.Max(f[k], f[k-amount*prices[i]]+amount*weight[i])
				}
			}
		}
	}
	return f[n]
}

//40ms 5.42MB
//拆分成根据 数量*金额的值 拆分成,01背包和完全背包
func backPackVII1(n int, prices []int, weight []int, amounts []int) int {
	//背包容量为i时,最大能买f[i]
	f := make([]int, n+1)
	m := len(prices)
	for i := 0; i < m; i++ {
		//分情况判断是否完全背包
		cost := prices[i] * amounts[i]
		if cost >= n {
			//n元钱,买不完i商品.完全背包
			for j := prices[i]; j <= n; j++ {
				f[j] = CommonUtil.Max(f[j], f[j-prices[i]]+weight[i])
			}
		} else {
			//01背包
			for j := n; j >= prices[i]; j-- {
				//多了个数量循环
				for k := 0; k <= amounts[i] && j-k*prices[i] >= 0; k++ {
					f[j] = CommonUtil.Max(f[j], f[j-k*prices[i]]+k*weight[i])
				}
			}
		}
	}
	return f[n]
}

//40ms 5.84MB
//转换成01背包
func backPackVII0(n int, prices []int, weight []int, amounts []int) int {
	//背包容量为i时,最大能买f[i]
	f := make([]int, n+1)
	m := len(prices)
	for i := 0; i < m; i++ {
		//01
		for j := n; j >= prices[i]; j-- {
			//多了个数量循环
			for k := 0; k <= amounts[i] && j-k*prices[i] >= 0; k++ {
				f[j] = CommonUtil.Max(f[j], f[j-k*prices[i]]+k*weight[i])
			}
		}
	}
	return f[n]
}
