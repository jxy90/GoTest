package middle_test

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_checkValidString(t *testing.T) {
	fmt.Println(checkValidString("()"))
}

func checkValidString(s string) bool {
	cntl, cntr := 0, 0
	for _, char := range s {
		switch char {
		case '(':
			cntl++
			cntr++
		case ')':
			cntl--
			cntr--
		case '*':
			cntl--
			cntr++
		}
		cntl = CommonUtil.Max(cntl, 0)
		if cntl > cntr {
			return false
		}
	}
	return cntl == 0
}
