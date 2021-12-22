package middle_test

import (
	"fmt"
	"strings"
	"testing"
)

func Test_repeatedStringMatch(t *testing.T) {
	fmt.Println(repeatedStringMatch("abcd", "cdabcdab"))
	fmt.Println(repeatedStringMatch("abc", "cabcabca"))
	fmt.Println(repeatedStringMatch("a", "aa"))
	fmt.Println(repeatedStringMatch("a", "a"))
	fmt.Println(repeatedStringMatch("abc", "wxyz"))
	fmt.Println(repeatedStringMatch("aa", "a"))
}
func repeatedStringMatch(a string, b string) int {
	cnt := 0
	temp := ""
	for (len(b)/len(a))+2 >= cnt {
		if strings.Index(temp, b) != -1 {
			return cnt
		}
		temp += a
		cnt++
	}
	return -1
}
