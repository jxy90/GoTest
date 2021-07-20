package middle_test

import (
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	println(groupAnagrams([]string{"huh", "tit"}))
	println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	cache := map[int][]string{}
	for _, str := range strs {
		temp := 0
		for _, v := range str {
			temp += 10 * int(v-'a')
		}
		cache[temp] = append(cache[temp], str)
	}
	ans := make([][]string, 0)
	for _, strings := range cache {
		ans = append(ans, strings)
	}
	return ans
}

func groupAnagrams0(strs []string) [][]string {
	cache := map[[26]int][]string{}
	for _, str := range strs {
		temp := [26]int{}
		for _, v := range str {
			temp[v-'a'] += 1
		}
		cache[temp] = append(cache[temp], str)
	}
	ans := make([][]string, 0)
	for _, strings := range cache {
		ans = append(ans, strings)
	}
	return ans
}
