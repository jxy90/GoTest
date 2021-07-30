package easy_test

import (
	"testing"
)

func Test_titleToNumber(t *testing.T) {
	println(titleToNumber("A"))
	println(titleToNumber("AB"))
	println(titleToNumber("ZY"))
	println(titleToNumber("FXSHRXW"))
}

func titleToNumber(columnTitle string) int {
	n := len(columnTitle)
	ans := 0
	for i := 0; i < n; i++ {
		ans = ans*26 + int(columnTitle[i]-'A'+1)
	}
	return ans
}
