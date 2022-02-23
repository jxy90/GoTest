package easy_test

import (
	"fmt"
	"testing"
)

func Test_reverseOnlyLetters(t *testing.T) {
	fmt.Println(reverseOnlyLetters("ab-cd"))
	fmt.Println(reverseOnlyLetters("7_28]"))
}

func reverseOnlyLetters(s string) string {
	ans := []byte(s)
	n := len(ans)
	left, right := 0, n-1
	for left < right {
		for !reverseOnlyLetters_isLetter(ans[left]) && left < right {
			left++
		}
		for !reverseOnlyLetters_isLetter(ans[right]) && left < right {
			right--
		}
		if left < right {
			ans[left], ans[right] = ans[right], ans[left]
			left++
			right--
		}
	}
	return string(ans)
}

func reverseOnlyLetters_isLetter(b byte) bool {
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
		return true
	}
	return false
}
