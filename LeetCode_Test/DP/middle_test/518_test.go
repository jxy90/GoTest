package middle_test

import (
	"fmt"
	"testing"
)

func Test_change(t *testing.T) {
	fmt.Println(change(5, []int{1, 2, 5}))
}

func change(amount int, coins []int) int {
	//无限个
	n := len(coins)
	f := make([]int, amount+1)
	f[0] = 1
	//f[i][j] 前i个硬币,组成j的方法数
	for i := 1; i <= n; i++ {
		a := coins[i-1]
		for j := a; j <= amount; j++ {
			f[j] += f[j-a]
		}
	}
	return f[amount]
}
func change0(amount int, coins []int) int {
	//无限个
	n := len(coins)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, amount+1)
		f[i][0] = 1
	}
	//f[i][j] 前i个硬币,组成j的方法数
	for i := 1; i <= n; i++ {
		a := coins[i-1]
		for j := 0; j <= amount; j++ {
			f[i][j] = f[i-1][j]
			if j-a >= 0 {
				f[i][j] += f[i][j-a]
			}
		}
	}
	return f[n][amount]
}
