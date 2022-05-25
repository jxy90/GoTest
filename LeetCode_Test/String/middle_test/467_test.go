package middle_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"strings"
	"testing"
)

func Test_467(t *testing.T) {
	t.Log(findSubstringInWraproundString("a"))
	t.Log(findSubstringInWraproundString("cac"))
	t.Log(findSubstringInWraproundString("zab"))
}

func findSubstringInWraproundString(p string) int {
	ans := 0
	f := make([]int, 26)
	j := 1
	for i := range p {
		if i > 0 && (p[i]-p[i-1]+26)%26 == 1 {
			j++
		} else {
			j = 1
		}
		f[p[i]-'a'] = CommonUtil.Max(f[p[i]-'a'], j)
	}
	for _, v := range f {
		ans += v
	}
	return ans
}
func findSubstringInWraproundString0(p string) int {
	strs := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
	n := len(p)
	cache := map[string]bool{}
	ans := 0
	for length := 1; length <= n; length++ {
		for start := 0; start+length <= n; start++ {
			sub := p[start : start+length]
			if !cache[sub] && strings.Contains(strs, sub) {
				ans++
				cache[sub] = true
			}
		}
	}
	return ans
}
