package Dynamic

import (
	"fmt"
	"strings"
	"testing"
)

//139. 单词拆分
//https://leetcode.cn/problems/word-break/description/

func wordBreak(s string, wordDict []string) bool {
	found := false
	n := len(wordDict)
	var helper func(start int)
	helper = func(start int) {
		if found {
			return
		}
		for i := 0; i < n; i++ {
			word := wordDict[i]
			if start+len(word) <= len(s) && s[start:start+len(word)] == word {
				if start+len(word) == len(s) {
					found = true
					return
				}
				helper(start + len(word))
			}
		}
	}
	helper(0)
	return found
}

//140. 单词拆分 II
//https://leetcode.cn/problems/word-break-ii/description/
func wordBreak0(s string, wordDict []string) []string {
	res := make([]string, 0)
	track := make([]string, 0)

	var helper func(start int)
	helper = func(start int) {
		if start == len(s) {
			res = append(res, strings.Join(track, " "))
			return
		}

		for i := 0; i < len(wordDict); i++ {
			word := wordDict[i]
			if start+len(word) <= len(s) && s[start:start+len(word)] == word {
				track = append(track, word)
				helper(start + len(word))
				track = track[:len(track)-1]
			}
		}
	}

	helper(0)
	return res

}

func Test_5(t *testing.T) {
	//fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak0("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}
