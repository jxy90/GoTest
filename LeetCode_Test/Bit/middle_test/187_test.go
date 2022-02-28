package middle_test

import (
	"fmt"
	"testing"
)

func Test_findRepeatedDnaSequences(t *testing.T) {
	//fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAA"))
}

//func findRepeatedDnaSequences(s string) []string {
//
//}

func findRepeatedDnaSequences(s string) []string {
	n := len(s)
	ans := make([]string, 0)
	if n < 10 {
		return ans
	}
	visited := make(map[string]int)
	for i := 0; i < n-10; i++ {
		if visited[s[i:i+10]] == 1 {
			ans = append(ans, s[i:i+10])
		}
		visited[s[i:i+10]]++
	}
	return ans
}
