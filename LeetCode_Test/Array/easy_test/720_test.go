package easy_test

import (
	"log"
	"sort"
	"testing"
)

func Test_720(t *testing.T) {
	//fmt.Println("123")
	log.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
}

func longestWord(words []string) (ans string) {
	sort.Slice(words, func(i, j int) bool {
		x, y := words[i], words[j]
		if len(x) < len(y) {
			return true
		} else if len(x) == len(y) {
			return x > y
		}
		return false
	})
	cache := make(map[string]bool, len(words)+1)
	cache[""] = true
	max := 0
	for _, word := range words {
		if _, ok := cache[word[:len(word)-1]]; ok && len(word) > max {
			cache[word] = true
			ans = word
		}
	}
	return ans
}
