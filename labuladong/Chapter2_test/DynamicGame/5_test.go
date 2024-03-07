package BackPack

import (
	"fmt"
	"math"
	"testing"
)

//887. 鸡蛋掉落
func superEggDrop(k int, n int) int {
	memo := make([][]int, k+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}
	var helper func(k int, n int) int
	helper = func(k int, n int) int {
		if k == 1 {
			return n
		}
		if n == 0 {
			return 0
		}
		if memo[k][n] != 0 {
			return memo[k][n]
		}
		res := math.MaxInt32
		for i := 1; i <= n; i++ {
			res = min(res, max(helper(k-1, i-1), helper(k, n-i))+1)
		}
		memo[k][n] = res
		return res
	}
	return helper(k, n)
}
func Test_5(t *testing.T) {
	fmt.Println(superEggDrop(2, 6))
	fmt.Println(superEggDrop(3, 14))
}
