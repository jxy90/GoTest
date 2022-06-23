package hard_test

import (
	"fmt"
	"testing"
)

//"wordgoodgoodgoodbestword"
//["word","good","best","good"]
func Test_30(t *testing.T) {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}
func findSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	m := len(words)
	length := m * wordLen

	cache := map[string]int{}
	for _, word := range words {
		cache[word]++
	}

	ans := make([]int, 0)
	cache2 := map[string]int{}
	cnt := 0
	for i := 0; i+length <= len(s); i++ {
		sub := s[i : i+length]
		for j := 0; j < m; j++ {
			item := sub[j*wordLen : (j+1)*wordLen]
			if _, ok := cache[item]; !ok {
				break
			} else {
				if cache2[item] == cache[item] {
					break
				}
				cache2[item]++
				cnt++
			}
		}
		if len(cache2) == len(cache) && cnt == m {
			ans = append(ans, i)
		}
		cache2 = map[string]int{}
		cnt = 0
	}
	return ans
}
