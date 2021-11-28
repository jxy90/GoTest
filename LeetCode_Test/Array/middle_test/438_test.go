package middle_test

import (
	"fmt"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}

func findAnagrams(s string, p string) (ans []int) {
	need, window := map[byte]int{}, map[byte]int{}
	for i := range p {
		need[p[i]]++
	}
	left, right := 0, 0
	valid := 0
	for right < len(s) {
		curR := s[right]
		right++
		if _, ok := need[curR]; ok {
			window[curR]++
			if window[curR] == need[curR] {
				valid++
			}
		}
		for right-left == len(p) {
			if valid == len(need) {
				ans = append(ans, left)
			}
			curL := s[left]
			left++
			if _, ok := need[curL]; ok {
				if window[curL] == need[curL] {
					valid--
				}
				window[curL]--
			}
		}
	}
	return
}
