package Array

import (
	"fmt"
	"testing"
)

//一道数组去重的算法题把我整不会了

//不同字符的最小子序列
func smallestSubsequence(s string) string {
	n := len(s)
	cnts := make(map[byte]int, n)
	for i := range s {
		cnts[s[i]]++
	}

	inStack := make(map[byte]bool, n)
	stack := make([]byte, 0, n)
	for i := range s {
		item := s[i]
		cnts[item]--

		if inStack[item] {
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > item && cnts[stack[len(stack)-1]] != 0 {
			inStack[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, item)
		inStack[item] = true
	}
	return string(stack)
}

func Test_10(t *testing.T) {
	fmt.Println(smallestSubsequence("cbacdcbc"))
}
