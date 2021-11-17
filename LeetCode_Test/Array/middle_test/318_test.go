package middle_test

import (
	"fmt"
	"testing"
)

func Test_maxProduct(t *testing.T) {
	fmt.Println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
}

func maxProduct(words []string) int {
	wordBits := make([]int, len(words))
	for i, word := range words {
		for _, b := range word {
			wordBits[i] |= 1 << (b - 'a')
		}
	}
	ans := 0
	for i := range words {
		for j := range words[:i] {
			if wordBits[i]&wordBits[j] == 0 && ans < len(words[i])*len(words[j]) {
				ans = len(words[i]) * len(words[j])
			}
		}
	}
	return ans
}
