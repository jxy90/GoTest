package hard_test

import (
	"fmt"
	"strings"
	"testing"
)

func isMatch(s string, p string) bool {
	slen := len(s)
	plen := len(p)
	dp := make([][]bool, slen+1)
	for i := range dp {
		dp[i] = make([]bool, plen+1)
	}
	for i := range p {
		if p[i] == '*' {
			dp[0][i+1] = true
		} else {
			break
		}
	}
	dp[0][0] = true
	for i := 0; i < slen; i++ {
		for j := 0; j < plen; j++ {
			if s[i] == p[j] || p[j] == '?' {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				dp[i+1][j+1] = dp[i+1][j] || dp[i][j+1]
			} else {
				dp[i+1][j+1] = false
			}
		}
	}
	return dp[slen][plen]
}
func isMatch1(s string, p string) bool {
	for strings.Contains(p, "**") {
		p = strings.ReplaceAll(p, "**", "*")
	}
	return isMatchHelper(s, p, 0, 0)
}

func isMatchHelper(s string, p string, i, j int) bool {
	//两个指针到指向终点 匹配
	if j == len(p) && i == len(s) {
		return true
	}
	//p剩余* 匹配
	if p[j:] == "*" {
		return true
	}
	//s指向终点 p还有剩余
	if i == len(s) {
		return false
	}
	// s和p剩余长度不相等，且p不包含*
	if !strings.Contains(p[j:], "*") && len(s[i:]) != len(p[j:]) {
		return false
	}
	if p[j] == s[i] {
		//相等，对比下一位
		return isMatchHelper(s, p, i+1, j+1)
	} else if p[j] == '?' {
		//p==？，对比下一位
		return isMatchHelper(s, p, i+1, j+1)
	} else if p[j] == '*' {
		//p==*。对剩余长度进行全匹配，党有一种情况满足，返回true
		flag := false
		for k := 0; len(s)-i-k > 0; k++ {
			flag = isMatchHelper(s, p, i+k, j+1)
			if flag {
				return flag
			}
		}
	}
	return false
}

func Test_isMath(t *testing.T) {
	fmt.Println(isMatch("a", ""))
	fmt.Println(isMatch("aaaa", "***a"))
	fmt.Println(isMatch("adceb", "****a*b"))
	fmt.Println(isMatch("", "******"))
	fmt.Println(isMatch("leetcode", "*e*t?d*"))
}
