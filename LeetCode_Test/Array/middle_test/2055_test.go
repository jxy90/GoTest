package middle_test

import (
	"fmt"
	"testing"
)

func Test_2055(t *testing.T) {
	fmt.Println(platesBetweenCandles("**|**|***|", [][]int{{2, 5}, {5, 9}}))
}

func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	//辅助数组
	g := make([]int, n)
	count := 0
	for i, b := range s {
		if b == '*' {
			count++
		} else {
		}
		g[i] = count
	}

	m := len(queries)
	ans := make([]int, m)
	for i, query := range queries {
		l, r := query[0], query[1]
		for s[l] != '|' {
			l++
			if l > r {
				break
			}
		}
		for s[r] != '|' {
			r--
			if l > r {
				break
			}
		}
		if l < r {
			ans[i] = g[r] - g[l]
		}
	}
	return ans
}
