package middle_test

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"math"
	"testing"
)

func Test_getMoneyAmount(t *testing.T) {
	fmt.Println(getMoneyAmount(10))
}

func getMoneyAmount(n int) int {
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			minCost := math.MaxInt32
			for k := i; k < j; k++ {
				cost := k + CommonUtil.Max(f[i][k-1], f[k+1][j])
				if minCost > cost {
					minCost = cost
				}
			}
			f[i][j] = minCost
		}
	}
	return f[1][n]
}
