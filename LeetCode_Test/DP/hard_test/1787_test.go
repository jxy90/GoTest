package DP_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"math"
	"testing"
)

func Test_minChanges(t *testing.T) {
	println(minChanges([]int{1, 2, 0, 3, 0}, 1))
}

func minChanges(nums []int, k int) int {
	n := len(nums)
	max := 1024
	f := make([][]int, k)
	for i := range f {
		f[i] = make([]int, max)
		for j := range f[i] {
			f[i][j] = math.MaxInt16
		}
	}
	g := make([]int, k)
	for i := range g {
		g[i] = math.MaxInt16
	}
	for i, cnt := 0, 0; i < k; i++ {
		// 使用 map 和 cnt 分别统计当前列的「每个数的出现次数」和「有多少个数」
		cache := make(map[int]int)
		for j := i; j < n; j += k {
			cache[nums[j]] += 1
			cnt++
		}
		if i == 0 {
			// 第 0 列：只需要考虑如何将该列变为 xor 即可
			for xor := 0; xor < max; xor++ {
				f[0][xor] = CommonUtil.Min(f[0][xor], cnt-cache[xor])
				g[0] = CommonUtil.Min(g[0], f[0][xor])
			}
		} else {
			// 其他列：考虑与前面列的关系
			for xor := 0; xor < max; xor++ {
				f[i][xor] = g[i-1] + cnt
				for cur := range cache {
					f[i][xor] = CommonUtil.Min(f[i][xor], f[i-1][xor^cur]+cnt-cache[cur])
				}
				g[i] = CommonUtil.Min(g[i], f[i][xor])
			}
		}

		cnt = 0
	}
	return f[k-1][0]
}
