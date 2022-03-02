package hard_test

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_numWays(t *testing.T) {
	fmt.Println(numWays(3, 2))
	fmt.Println(numWays(2, 4))
	fmt.Println(numWays(4, 2))
	fmt.Println(numWays(47, 38))
}

var mod = 1000000000 + 7

func numWays(steps int, arrLen int) int {
	arrLen = CommonUtil.Min(steps, arrLen)
	f := make([]int, arrLen+1)

	options := []int{-1, 0, 1}
	//f[i][j] 表示通过i步到j位置的方案数目
	f[0] = 1
	for i := 0; i < steps; i++ {
		nextf := make([]int, arrLen+1)
		for j := 0; j < arrLen; j++ {
			for k := 0; k < 3; k++ {
				nextJ := j + options[k]
				if nextJ >= 0 && nextJ < arrLen {
					nextf[nextJ] = (nextf[nextJ] + f[j]) % mod
				}
			}
		}
		f = nextf
	}
	return f[0]
}

func numWays0(steps int, arrLen int) int {
	f := make([][]int, steps+1)
	for i := range f {
		f[i] = make([]int, arrLen+1)
	}
	options := []int{-1, 0, 1}
	//f[i][j] 表示通过i步到j位置的方案数目
	f[0][0] = 1
	for i := 0; i < steps; i++ {
		for j := 0; j < arrLen; j++ {
			for k := 0; k < 3; k++ {
				nextJ := j + options[k]
				if nextJ >= 0 && nextJ < arrLen {
					f[i+1][nextJ] = (f[i+1][nextJ] + f[i][j]) % mod
				}
			}
		}
	}
	return f[steps][0]
}
