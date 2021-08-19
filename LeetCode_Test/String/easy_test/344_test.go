package easy_test

import "testing"

func Test_reverseString(t *testing.T) {
	reverseString([]byte{'1'})
}
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
