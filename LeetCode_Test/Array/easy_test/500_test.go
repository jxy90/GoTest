package easy_test

import (
	"fmt"
	"testing"
)

func Test_findWords(t *testing.T) {
	fmt.Println('A')
	fmt.Println('a')
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}

func findWords(words []string) []string {
	keys := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	cache := map[byte]int{}
	for i, key := range keys {
		for j := range key {
			cache[key[j]] = i
			cache[key[j]-32] = i
		}
	}
	ans := make([]string, 0)
	for _, word := range words {
		flag := true
		for i := range word {
			if cache[word[i]] != cache[word[0]] {
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, word)
		}
	}
	return ans
}
