package Chapter0_test

import (
	"math"
	"testing"
)

func minWindow(s string, t string) string {
	need, window := map[byte]int{}, map[byte]int{}
	for i := range t {
		need[t[i]]++
	}
	left, right := 0, 0
	valid := 0
	index, length := 0, math.MaxInt32
	for right < len(s) {
		tempAdd := s[right]
		right++

		if _, ok := need[tempAdd]; ok {
			window[tempAdd]++
			if window[tempAdd] == need[tempAdd] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				index = left
				length = right - left
			}
			tempDel := s[left]
			left++

			if _, ok := need[tempDel]; ok {
				window[tempDel]--
				if window[tempDel] < need[tempDel] {
					valid--
				}
			}
		}

	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[index : index+length]
}

func checkInclusion(s1 string, s2 string) bool {
	need, window := map[byte]uint8{}, map[byte]uint8{}
	for i := range s1 {
		need[s1[i]]++
	}

	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		tempAdd := s2[right]
		right++

		if _, ok := need[tempAdd]; ok {
			window[tempAdd]++
			if window[tempAdd] == need[tempAdd] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			tempDel := s2[left]
			left++

			if _, ok := need[tempDel]; ok {
				if window[tempDel] == need[tempDel] {
					valid--
				}
				window[tempDel]--
			}
		}
	}
	return false
}

func findAnagrams(s string, p string) []int {
	need, window := map[byte]int{}, map[byte]int{}
	for i := range p {
		need[p[i]]++
	}
	left, right := 0, 0
	valid := 0
	indexes := []int{}
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left >= len(p) {
			if valid == len(need) {
				indexes = append(indexes, left)
			}
			d := s[left]
			left++
			if _, ok := window[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return indexes
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	window := map[byte]int{}
	left, right := 0, 0
	length := 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		if left == 0 {
			length = right - 1
		}

		for window[c] != 1 {
			d := s[left]
			left++
			window[d]--
		}
		if right-left > length {
			length = right - left
		}
	}
	return length
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	println(lengthOfLongestSubstring(" "))
}
