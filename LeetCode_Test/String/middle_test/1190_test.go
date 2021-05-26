package middle_test

import (
	"strings"
	"testing"
)

func Test_reverseParentheses(t *testing.T) {
	println(reverseParentheses("(abcd)"))
	println(reverseParentheses("(u(love)i)"))
	println(reverseParentheses("(ed(et(oc))el)"))
}

func reverseParentheses(s string) string {
	stack := make([]byte, 0)
	for i := range s {
		stack = append(stack, s[i])
		//遇到)取出到上一个(,再反转插入
		if s[i] == ')' {
			//取出上一组()所有数
			temp := make([]byte, 0)
			for stack[len(stack)-1] != '(' {
				stack = stack[:len(stack)-1]
				temp = append(temp, stack[len(stack)-1])
			}
			//删除(
			stack = stack[:len(stack)-1]
			temp = temp[:len(temp)-1]
			//反转插入stack
			stack = append(stack, temp...)
		}
	}
	sb := strings.Builder{}
	sb.Write(stack)
	return sb.String()
}
