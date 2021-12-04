package easy_test

import (
	"fmt"
	"testing"
)

func Test_canConstruct(t *testing.T) {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}

func canConstruct(ransomNote string, magazine string) bool {
	cache := map[byte]int{}
	for i, _ := range magazine {
		//fmt.Println(magazine[i])
		//fmt.Println(v)
		cache[magazine[i]]++
	}
	for i := range ransomNote {
		cache[ransomNote[i]]--
		if cache[ransomNote[i]] < 0 {
			return false
		}
	}
	return true
}
