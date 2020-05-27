package easy_test

import "testing"

func reverseStr(s string, k int) string {
	sSlice := []byte(s)
	for i := 0; i < len(s); i = i + 2*k {
		left := i
		right := min(i+k-1, len(s)-1)
		for left < right {
			sSlice[left], sSlice[right] = sSlice[right], sSlice[left]
			left++
			right--
		}
	}
	return string(sSlice)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Test_reverseStr(t *testing.T) {
	println(reverseStr("abcdef", 3))
}
