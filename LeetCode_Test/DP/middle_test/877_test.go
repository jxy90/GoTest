package middle

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_stoneGame(t *testing.T) {
	fmt.Println(stoneGame([]int{5, 3, 4, 5}))
}

func stoneGame(piles []int) bool {
	//分析亚先手,李后手.
	//堆数为偶数,总石子数为奇数,只能从头部或者尾部拿
	//f[i][j]表示第i颗石子到第j颗石子,先后手的差值
	n := len(piles)
	f := make([][]int, n+2)
	for i := range f {
		f[i] = make([]int, n+2)
	}
	for length := 1; length <= n; length++ {
		for l := 1; l+length-1 <= n; l++ {
			r := l + length - 1
			a := piles[l-1] - f[l+1][r]
			b := piles[r-1] - f[l][r-1]
			f[l][r] = CommonUtil.Max(a, b)
		}
	}
	return f[1][n] > 0
}

func stoneGame0(piles []int) bool {
	//分析亚先手,李后手.
	//堆数为偶数,总石子数为奇数,只能从头部或者尾部拿
	//每次都是先手从两边调最大的
	return true
}
