package easy_test

import "testing"

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func Test_reverseString(t *testing.T) {
	reverseString([]byte{'1'})
}
