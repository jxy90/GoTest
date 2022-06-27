package middle_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_522(t *testing.T) {
}

func findLUSlength(strs []string) int {
	ans := -1
	for i, s := range strs {
		check := true
		for j, t := range strs {
			if i != j && isSubSeq(s, t) {
				check = false
				break
			}
		}
		if check {
			ans = CommonUtil.Max(ans, len(s))
		}
	}
	return ans
}

func isSubSeq(s, t string) bool {
	p1 := 0
	for i := range t {
		if t[i] == s[p1] {
			p1++
			if p1 == len(s) {
				return true
			}
		}
	}
	return false
}
