package middle_test

import (
	"fmt"
	"testing"
)

func Test_permutation(t *testing.T) {
	fmt.Println(permutation("abc"))
}

func permutation(s string) []string {
	//判断当前结果是否存在ans中
	cache := map[string]bool{}
	//判断字符是否被访问
	visited := map[int]bool{}

	n := len(s)
	var dfs func(str string)
	dfs = func(str string) {
		if n == len(str) {
			cache[str] = true
			return
		}
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			dfs(str + string(s[i]))
			visited[i] = false
		}
	}
	dfs("")
	ans := make([]string, 0)
	for s := range cache {
		ans = append(ans, s)
	}
	return ans
}
