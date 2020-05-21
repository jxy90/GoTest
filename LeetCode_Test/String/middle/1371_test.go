package middle_test

import (
	"strings"
	"testing"
)

var ans int

func findTheLongestSubstring(s string) int {
	ans = 0
	length := len(s)
	set := map[int32]int{}
	var i, j int
	for i = 0; i < length; i++ {
		for j = i + 1; j <= length; j++ {
			if j-i>ans {
				findTheLongestSubstringInitSet(set)
				if findTheLongestSubstringCheck(s[i:j], set) {
						ans = j - i
				}
			}else {
				j = i+ans
			}
		}
	}
	return ans
}

func findTheLongestSubstringInitSet(set map[int32]int) {
	set['a'] = 0
	set['e'] = 0
	set['i'] = 0
	set['o'] = 0
	set['u'] = 0
}
func findTheLongestSubstringCheck(substr string, set map[int32]int) bool {
	for _, item := range substr {
		itemStr := string(item)
		if strings.Contains("aeiou", itemStr) {
			set[item] += 1
		}
	}
	for _, item := range set {
		if item%2 == 1 {
			return false
		}
	}
	return true
}

func Test_findTheLongestSubstring(t *testing.T) {
	str := "id"
	findTheLongestSubstring(str)


	pos := make([]int, 1 << 4)
	println(pos)

}
