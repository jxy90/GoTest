package middle_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"math"
	"testing"
)

func Test_numSquares(t *testing.T) {
	println(numSquares(12))
	println(numSquares(13))
}

func numSquares(n int) int {
	//n=12 count=3 得到可以使用的数字
	count := int(math.Floor(math.Sqrt(float64(n))))
	f := make([]int, n+1)
	for i := range f {
		f[i] = i
	}
	for i := 0; i <= count; i++ {
		for j := i * i; j <= n; j++ {
			f[j] = CommonUtil.Min(f[j], f[j-i*i]+1)
		}
	}
	return f[n]
}
