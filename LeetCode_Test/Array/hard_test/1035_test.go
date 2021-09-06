package hard_test

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_maxUncrossedLines(t *testing.T) {
	//fmt.Println(maxUncrossedLines([]int{1, 4, 2}, []int{1, 2, 4}))
	//fmt.Println(maxUncrossedLines([]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}))
	fmt.Println(maxUncrossedLines([]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}))
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	//f[i][j] 第i,j个位置连线的方案数
	m := len(nums1)
	n := len(nums2)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			f[i][j] = CommonUtil.Max(f[i-1][j], f[i][j-1])
			if nums1[i] == nums2[j] {
				f[i][j] = CommonUtil.Max(f[i-1][j-1]+1, f[i][j])
			}
		}
	}
	return f[m][n]
}
