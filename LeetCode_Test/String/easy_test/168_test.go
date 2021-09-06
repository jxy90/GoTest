package easy_test

import (
	"fmt"
	"testing"
)

func Test_convertToTitle(t *testing.T) {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(26))
	fmt.Println(convertToTitle(27))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))
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
