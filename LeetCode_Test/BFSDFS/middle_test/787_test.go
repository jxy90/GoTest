package middle_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"math"
	"testing"
)

func Test_findCheapestPrice(t *testing.T) {

	fmt.Println(findCheapestPrice(5, [][]int{{4, 1, 1}, {1, 2, 3}, {0, 3, 2}, {0, 4, 10}, {3, 1, 1}, {1, 4, 3}}, 2, 1, 1))
	//fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1))
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	f := make([][]int, k+2)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = math.MaxInt32
		}
	}
	f[0][src] = 0
	for t := 0; t < k+2; t++ {
		for _, flight := range flights {
			i, j, cost := flight[0], flight[1], flight[2]
			f[t][j] = CommonUtil.Min(f[t][j], f[t-1][i]+cost)
		}
	}
	ans := math.MaxInt32
	for t := 0; t < k+2; t++ {
		ans = CommonUtil.Min(ans, f[t][dst])
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
func findCheapestPrice0(n int, flights [][]int, src int, dst int, k int) int {
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = math.MaxInt32
		}
	}
	for _, flight := range flights {
		a, b, c := flight[0], flight[1], flight[2]
		f[a][b] = c
	}
	ans := math.MaxInt32
	var dfs func(s, d, count, totalPrice int)
	dfs = func(s, d, count, totalPrice int) {
		if count > k+1 || totalPrice >= ans {
			return
		}
		if s == dst {
			ans = CommonUtil.Min(ans, totalPrice)
			return
		}
		for i, price := range f[s] {
			if price == math.MaxInt32 {
				continue
			}
			dfs(i, d, count+1, totalPrice+price)
		}
	}
	dfs(src, dst, 0, 0)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
