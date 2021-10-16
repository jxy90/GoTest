package hard_test

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_addOperators(t *testing.T) {
	fmt.Println(isMatch("a", ""))
	fmt.Println(isMatch("aaaa", "***a"))
	fmt.Println(isMatch("adceb", "****a*b"))
	fmt.Println(isMatch("", "******"))
	fmt.Println(isMatch("leetcode", "*e*t?d*"))
}

var n, t int
var ans []string
var s string

func addOperators(num string, target int) []string {
	n = len(num)
	t = target
	ans = make([]string, 0)
	s = num
	addOperatorsDFS(0, 0, 0, "")
	return ans
}

func addOperatorsDFS(index, prev, value int, str string) {
	if index == n {
		if value == t {
			ans = append(ans, str)
		}
		return
	}
	for i := index; i < n; i++ {
		if index != i && s[index] == '0' {
			break
		}
		next, _ := strconv.Atoi(s[index : i+1])
		nextStr := s[index : i+1]
		if index == 0 {
			addOperatorsDFS(i+1, next, next, ""+nextStr)
		} else {
			addOperatorsDFS(i+1, next, value+next, str+"+"+nextStr)
			addOperatorsDFS(i+1, -next, value-next, str+"-"+nextStr)
			x := prev * next
			addOperatorsDFS(i+1, x, value-prev+x, str+"*"+nextStr)
		}
	}
}
