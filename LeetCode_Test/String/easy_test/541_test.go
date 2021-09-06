package easy_test

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_reverseStr(t *testing.T) {
	fmt.Println(reverseStr("abcdef", 2))
}

func reverseStr(s string, k int) string {
	sBytes := []byte(s)
	n := len(s)
	for i := 0; i < n; i = i + 2*k {
		l, r := i, CommonUtil.Min(i+k-1, n-1)
		for l < r {
			sBytes[l], sBytes[r] = sBytes[r], sBytes[l]
			l++
			r--
		}
	}
	return string(sBytes)
}
