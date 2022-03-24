package easy_test

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_661(t *testing.T) {
	//{{1,1,1},{1,0,1},{1,1,1}}
	//
	//{{100,200,100},{200,50,200},{100,200,100}}
	fmt.Println(imageSmoother([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
	fmt.Println(imageSmoother([][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}))
}

func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	ans := make([][]int, m)
	g := make([][]int, m+1)
	for i := range ans {
		ans[i] = make([]int, n)
		g[i] = make([]int, n+1)
	}
	g[m] = make([]int, n+1)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			g[i][j] = img[i-1][j-1] + g[i][j-1] + g[i-1][j] - g[i-1][j-1]
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			startX := CommonUtil.Max(0, i-1)
			endX := CommonUtil.Min(i+1, m-1)
			startY := CommonUtil.Max(0, j-1)
			endY := CommonUtil.Min(j+1, n-1)
			ans[i][j] = (g[endX+1][endY+1] - g[endX+1][startY] - g[startX][endY+1] + g[startX][startY]) / ((endX + 1 - startX) * (endY + 1 - startY))
		}
	}
	return ans
}
