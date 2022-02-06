package hard_test

import (
	"fmt"
	"strings"
	"testing"
)

func Test_longestDupSubstring(t *testing.T) {
	fmt.Println(longestDupSubstring("banana"))
	fmt.Println(longestDupSubstring("abcd"))
}

func longestDupSubstring(s string) string {
	n := len(s)
	l, r := 0, n-1
	ans := ""
	for l < r {
		mid := l + (r-l+1)>>1
		if res, ok := longestDupSubstringCheck(s, mid); ok {
			ans = res
			l = mid
		} else {
			r = mid - 1
		}
	}
	return ans
}

func longestDupSubstringCheck(s string, length int) (string, bool) {
	cache := map[string]struct{}{}
	for i := 0; i <= len(s)-length; i++ {
		subStr := s[i : i+length]
		if _, ok := cache[subStr]; ok {
			return subStr, true
		}
		cache[subStr] = struct{}{}
	}
	return "", false
}

func longestDupSubstringCheck0(s string, length int) (string, bool) {
	for i := 0; i < len(s)-length; i++ {
		index0 := strings.Index(s, s[i:i+length])
		index1 := strings.LastIndex(s, s[i:i+length])
		if index0 != index1 {
			return s[i : i+length], true
		}
	}
	return "", false
}
