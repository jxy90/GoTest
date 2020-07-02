package easy_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func reverseStr(s string, k int) string {
	sSlice := []byte(s)
	for i := 0; i < len(s); i = i + 2*k {
		left := i
		right := CommonUtil.Min(i+k-1, len(s)-1)
		for left < right {
			sSlice[left], sSlice[right] = sSlice[right], sSlice[left]
			left++
			right--
		}
	}
	return string(sSlice)
}

func Test_reverseStr(t *testing.T) {
	println(reverseStr("abcdef", 3))
}
