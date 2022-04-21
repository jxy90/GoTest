package easy_test

import (
	"log"
	"strings"
	"testing"
)

func Test_824(t *testing.T) {
	log.Println(toGoatLatin("I speak Goat Latin"))
}

func toGoatLatin(sentence string) string {
	yuanYin := map[byte]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}, 'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {}}
	n := len(sentence)
	wordCnt := 0
	ans := strings.Builder{}
	for start := 0; start < n; start++ {
		if wordCnt > 0 {
			ans.WriteByte(' ')
		}
		end := start
		for ; end < n && sentence[end] != ' '; end++ {
		}
		if _, ok := yuanYin[sentence[start]]; ok {
			ans.WriteString(sentence[start+1 : end])
			ans.WriteByte(sentence[start])
		} else {
			ans.WriteString(sentence[start:end])
		}
		ans.WriteByte('m')
		ans.WriteString(strings.Repeat("a", wordCnt+2))
		start = end
		wordCnt++
	}
	return ans.String()
}
