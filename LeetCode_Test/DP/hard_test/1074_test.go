package DP_test

import (
	"testing"
)

func Test_numSubmatrixSumTarget(t *testing.T) {
	fmt.Println(numSubmatrixSumTarget([][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, 0))
	fmt.Println(numSubmatrixSumTarget([][]int{{1, -1}, {-1, 1}}, 0))
	fmt.Println(numSubmatrixSumTarget([][]int{{987}}, 0))
}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m := len(matrix)
	n := len(matrix[0])
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			f[i][j] = matrix[i-1][j-1] + f[i-1][j] + f[i][j-1] - f[i-1][j-1]
		}
	}
	ans := 0
	for top := 1; top <= m; top++ {
		for bot := top; bot <= m; bot++ {
			cur := 0
			cache := make(map[int]int)
			for r := 1; r <= n; r++ {
				cur = f[bot][r] - f[top-1][r]
				if cur == target {
					ans++
				}
				if _, ok := cache[cur-target]; ok {
					ans += cache[cur-target]
				}
				cache[cur]++
			}
		}
	}
	return ans
}

func numSubmatrixSumTarget0(matrix [][]int, target int) int {
	m := len(matrix)
	n := len(matrix[0])
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			f[i][j] = matrix[i-1][j-1] + f[i-1][j] + f[i][j-1] - f[i-1][j-1]
		}
	}
	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for p := i; p <= m; p++ {
				for q := j; q <= n; q++ {
					if f[p][q]-f[i-1][q]-f[p][j-1]+f[i-1][j-1] == target {
						//fmt.Println(fmt.Sprintf("i:%v j:%v p:%v q:%v", i, j, p, q))
						ans++
					}
				}
			}
		}
	}
	return ans
}
