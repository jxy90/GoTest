package middle_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_longestDiverseString(t *testing.T) {
	fmt.Println(longestDiverseString(1, 1, 7))
	fmt.Println(longestDiverseString(2, 2, 1))
	fmt.Println(longestDiverseString(7, 1, 0))
}

type pair struct {
	b byte
	c int
}

func longestDiverseString(a int, b int, c int) (ans string) {
	pairs := []pair{{'a', a}, {'b', b}, {'c', c}}
	for {
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].c > pairs[j].c
		})
		for i := range pairs {
			if pairs[i].c == 0 {
				return
			}
			if len(ans) < 2 || ans[len(ans)-1] != ans[len(ans)-2] || ans[len(ans)-1] != pairs[i].b {
				ans += string(pairs[i].b)
				pairs[i].c--
				break
			}
		}
	}
	return
}
