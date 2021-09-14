package middle_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_findLongestWord(t *testing.T) {
	//fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea", "abpcplaaa", "abpcllllll", "abccclllpppeeaaaa"}))
}

func findLongestWord(s string, dictionary []string) string {
	sort.Strings(dictionary)
	ans := ""
	lengthS := len(s)
	for _, dic := range dictionary {
		lengthDic := len(dic)
		i, j := 0, 0
		for i < lengthS && j < lengthDic {
			if s[i] == dic[j] {
				i++
				j++
			} else {
				i++
			}
		}
		lengthAns := len(ans)
		if j == lengthDic && (lengthDic > lengthAns || (lengthDic == lengthAns && dic < ans)) {
			ans = dic
		}
	}
	return ans
}
