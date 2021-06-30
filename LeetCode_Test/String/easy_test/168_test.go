package easy_test

import (
	"testing"
)

func Test_convertToTitle(t *testing.T) {
	println(convertToTitle(1))
	println(convertToTitle(26))
	println(convertToTitle(27))
	println(convertToTitle(28))
	println(convertToTitle(701))
}

func convertToTitle(columnNumber int) string {
	ans := []byte{}
	for columnNumber > 0 {
		columnNumber--
		ans = append([]byte{'A' + byte(columnNumber%26)}, ans...)
		columnNumber /= 26
	}
	return string(ans)
}
