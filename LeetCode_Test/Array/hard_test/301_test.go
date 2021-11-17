package hard_test

import (
	"fmt"
	"testing"
)

func Test_removeInvalidParentheses(t *testing.T) {
	fmt.Println(removeInvalidParentheses("()())()"))
}

func removeInvalidParentheses(s string) (ans []string) {
	lr, rr := 0, 0
	for _, b := range s {
		if b == '(' {
			lr++
		} else if b == ')' {
			if lr == 0 {
				rr++
			} else {
				lr--
			}
		}
	}
	removeInvalidParenthesesHelper(&ans, s, 0, lr, rr)
	return
}

func removeInvalidParenthesesHelper(ans *[]string, s string, start, lr, rr int) {
	if lr == 0 && rr == 0 {
		if removeInvalidParenthesesValid(s) {
			*ans = append(*ans, s)
		}
		return
	}

	for i := start; i < len(s); i++ {
		if i != start && s[i] == s[i-1] {
			continue
		}
		//长度无法满足
		if lr+rr > len(s)-i {
			return
		}
		//去掉左括号
		if lr > 0 && s[i] == '(' {
			removeInvalidParenthesesHelper(ans, s[:i]+s[i+1:], i, lr-1, rr)
		}
		//去掉有括号
		if rr > 0 && s[i] == ')' {
			removeInvalidParenthesesHelper(ans, s[:i]+s[i+1:], i, lr, rr-1)
		}
	}
}

func removeInvalidParenthesesValid(str string) bool {
	cnt := 0
	for _, ch := range str {
		if ch == '(' {
			cnt++
		} else if ch == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}
