package Chapter1_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

//经典背包
func BeiBao(wt, val []int, N, W int) int {
	dp := make([][]int, N+1)
	for index := range dp {
		dp[index] = make([]int, W+1)
	}
	for i := 1; i <= N; i++ {
		for w := 1; w <= W; w++ {
			//println("N:", strconv.Itoa(i), " W:", strconv.Itoa(w))
			if w-wt[i-1] > 0 {
				dp[i][w] = CommonUtil.Max(dp[i-1][w], dp[i-1][w-wt[i-1]]+val[i-1])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}
	return dp[N][W]
}

func Test_BB(t *testing.T) {
	N, W := 3, 4
	wt := []int{2, 1, 3}
	val := []int{4, 2, 3}
	data := BeiBao(wt, val, N, W)
	println(data)
}
