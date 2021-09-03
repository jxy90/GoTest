package middle

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_lastStoneWeightII(t *testing.T) {
	fmt.Println(lastStoneWeightII([]int{1, 1, 4, 2, 2}))
	fmt.Println(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println(lastStoneWeightII([]int{31, 26, 33, 21, 40}))
	fmt.Println(lastStoneWeightII([]int{1, 2}))
}

func lastStoneWeightII(stones []int) int {
	// 把石头分成两堆要求两堆重量尽可能接近；定义一个容量为sum/2的01背包
	sum := 0
	for i := range stones {
		sum += stones[i]
	}
	n := len(stones)
	m := sum / 2
	//f[i][j] 代表考虑前 i 个物品（数值），凑成总和不超过 j 的最大价值
	f := make([]int, m+1)

	//f[0][0] = 1
	for i := 1; i <= n; i++ {
		x := stones[i-1]
		for j := m; j >= x; j-- {
			f[j] = CommonUtil.Max(f[j], f[j-x]+x)
		}
	}
	return CommonUtil.Abs(sum - f[m] - f[m])
}

func lastStoneWeightII0(stones []int) int {
	// 把石头分成两堆要求两堆重量尽可能接近；定义一个容量为sum/2的01背包
	sum := 0
	for i := range stones {
		sum += stones[i]
	}
	n := len(stones)
	m := sum / 2
	//f[i][j] 代表考虑前 i 个物品（数值），凑成总和不超过 j 的最大价值
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	//f[0][0] = 1
	for i := 1; i <= n; i++ {
		x := stones[i-1]
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j]
			if j >= x {
				f[i][j] = CommonUtil.Max(f[i][j], f[i-1][j-x]+x)
			}
		}
	}
	return CommonUtil.Abs(sum - f[n][m] - f[n][m])
}
