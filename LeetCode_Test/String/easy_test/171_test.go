package easy_test

import (
	"fmt"
	"testing"
)

func Test_titleToNumber(t *testing.T) {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))
	fmt.Println(titleToNumber("FXSHRXW"))
}

func titleToNumber(columnTitle string) int {
	n := len(columnTitle)
	ans := 0
	for i := 0; i < n; i++ {
		ans = ans*26 + int(columnTitle[i]-'A'+1)
	}
	return ans
}
