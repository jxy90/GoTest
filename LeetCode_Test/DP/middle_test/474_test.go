package middle_test

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_findMaxForm(t *testing.T) {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	fmt.Println(findMaxForm([]string{"10", "0", "1"}, 1, 1))
}

func findMaxForm(strs []string, m int, n int) int {
	length := len(strs)
	//f := make([][]int, length+1)
	//for i := range strs {
	//	f[i+1] = findMaxFormCheck(strs[i])
	//}
	g := make([][]int, m+1)
	for j := range g {
		g[j] = make([]int, n+1)
	}
	for i := 1; i <= length; i++ {
		zeroOne := findMaxFormCheck(strs[i-1])
		for j := m; j >= zeroOne[0]; j-- {
			for k := n; k >= zeroOne[1]; k-- {
				g[j][k] = CommonUtil.Max(g[j][k], g[j-zeroOne[0]][k-zeroOne[1]]+1)
			}
		}
	}
	return g[m][n]
}
func findMaxForm0(strs []string, m int, n int) int {
	length := len(strs)
	//f := make([][]int, length+1)
	//for i := range strs {
	//	f[i+1] = findMaxFormCheck(strs[i])
	//}
	g := make([][][]int, length+1)
	for i := range g {
		g[i] = make([][]int, m+1)
		for j := range g[i] {
			g[i][j] = make([]int, n+1)
		}
	}
	for i := 1; i <= length; i++ {
		zeroOne := findMaxFormCheck(strs[i-1])
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				g[i][j][k] = g[i-1][j][k]
				if j >= zeroOne[0] && k >= zeroOne[1] {
					g[i][j][k] = CommonUtil.Max(g[i][j][k], g[i-1][j-zeroOne[0]][k-zeroOne[1]]+1)
				}
			}
		}
	}
	return g[length][m][n]
}

func findMaxFormCheck(str string) []int {
	cntN := 0
	cntM := 0
	for i := range str {
		if str[i] == '0' {
			cntM++
		} else {
			cntN++
		}
	}
	return []int{cntM, cntN}
}
