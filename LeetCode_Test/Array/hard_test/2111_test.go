package hard_test

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_kIncreasing(t *testing.T) {
	//fmt.Println(kIncreasing([]int{5, 4, 3, 2, 1}, 1))
	fmt.Println(kIncreasing([]int{4, 1, 5, 2, 6, 2}, 2))
	fmt.Println(kIncreasing([]int{4, 1, 5, 2, 6, 2}, 3))
}

func kIncreasing(arr []int, k int) int {
	n := len(arr)
	ans := n
	f := make([]int, n)
	g := make([]int, k)
	for i := 0; i < k; i++ {
		g[i] = 1
		for j := i; j < n; j += k {
			f[j] = 1
			for p := i; p < j; p += k {
				if arr[j] >= arr[p] {
					f[j] = CommonUtil.Max(f[j], f[p]+1)
				}
			}
			g[i] = CommonUtil.Max(g[i], f[j])
		}
	}
	for _, i := range g {
		ans -= i
	}
	return ans
}
