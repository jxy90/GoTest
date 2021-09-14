package easy_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_longestWord(t *testing.T) {
	fmt.Println(longestWord([]string{"w", "wo", "wor", "worl", "world"}))
}

func longestWord(words []string) string {
	set := map[string]bool{}
	sort.Strings(words)
	//fmt.Println(words)
	ans := ""
	for _, word := range words {
		set[word] = true
		flag := true
		for i := 1; i < len(word); i++ {
			if !set[word[:i]] {
				flag = false
				break
			}
		}
		if flag && len(ans) < len(word) && (len(ans) == len(word) && ans > word) {
			ans = word
		}
	}
	return ans
}
