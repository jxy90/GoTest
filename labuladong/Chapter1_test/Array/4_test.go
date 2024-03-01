package Array

import (
	"math"
	"testing"
)

//滑动窗口

//最小覆盖子串
//https://leetcode.cn/problems/minimum-window-substring/description/
func minWindow(s string, t string) string {
	windows := make(map[byte]int)
	needs := make(map[byte]int)
	for i := range t {
		needs[t[i]]++
	}

	left, right := 0, 0
	start, length := 0, math.MaxInt32
	valid := 0
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := needs[c]; ok {
			windows[c]++
			if windows[c] == needs[c] {
				valid++
			}
		}

		for left < len(s) && valid == len(needs) {
			if right-left < length {
				start = left
				length = right - left
			}
			c = s[left]
			left++
			if _, ok := needs[c]; ok {
				if windows[c] == needs[c] {
					valid--
				}
				windows[c]--
			}
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}

//字符串的排列
//https://leetcode.cn/problems/permutation-in-string/
func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := range s1 {
		need[s1[i]]++
	}
	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		c := s2[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left == len(s1) {
				return true
			}
			c := s2[left]
			left++
			if _, ok := need[c]; ok {
				if window[c] == need[c] {
					valid--
				}
				window[c]--
			}
		}
	}
	return false
}

//找到字符串中所有字母异位词
//https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/
func findAnagrams(s string, p string) []int {
	windows := make(map[byte]int)
	need := make(map[byte]int)
	for i := range p {
		need[p[i]]++
	}
	left, right := 0, 0
	valid := 0
	//start, length := 0, math.MaxInt32
	ans := make([]int, 0)
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			windows[c]++
			if windows[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left == len(p) {
				ans = append(ans, left)
			}
			c := s[left]
			left++

			if _, ok := need[c]; ok {
				if windows[c] == need[c] {
					valid--
				}
				windows[c]--
			}
		}
	}
	return ans
}

//无重复字符的最长子串
//https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	windows := make(map[byte]int)
	length := 0
	for right < len(s) {
		c := s[right]
		right++
		windows[c]++
		for windows[c] > 1 {
			d := s[left]
			left++
			windows[d]--
		}
		if right-left > length {
			length = right - left
		}
	}
	return length
}

//https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
func dismantlingAction(arr string) int {
	length := 0
	left, right := 0, 0
	windows := make(map[byte]int)
	for right < len(arr) {
		c := arr[right]
		right++
		windows[c]++

		for windows[c] > 1 {
			d := arr[left]
			left++
			windows[d]--
		}
		if right-left > length {
			length = right - left
		}
	}
	return length
}

//https://leetcode.cn/problems/MPnaiL/description/
func checkInclusion2(s1 string, s2 string) bool {
	need := make(map[byte]int)
	windows := make(map[byte]int)
	for i := range s1 {
		need[s1[i]]++
	}
	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		c := s2[right]
		right++
		if _, ok := need[c]; ok {
			windows[c]++
			if windows[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left == len(s1) {
				return true
			}
			d := s2[left]
			left++
			if _, ok := need[d]; ok {
				if windows[c] == need[c] {
					valid--
				}
				windows[c]--
			}
		}
	}
	return false
}

//https://leetcode.cn/problems/VabMRr/
func findAnagrams2(s string, p string) []int {
	need := make(map[byte]int)
	windows := make(map[byte]int)
	for i := range p {
		need[p[i]]++
	}
	left, right := 0, 0
	valid := 0
	ans := make([]int, 0)
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			windows[c]++
			if windows[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left == len(p) {
				ans = append(ans, left)
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if windows[d] == need[d] {
					valid--
				}
				windows[d]--
			}
		}
	}
	return ans
}

//https://leetcode.cn/problems/wtcaE1/
func lengthOfLongestSubstring1(s string) int {
	windows := make(map[byte]int)
	left, right := 0, 0
	length := 0
	for right < len(s) {
		c := s[right]
		right++
		windows[c]++

		for windows[c] > 1 {
			d := s[left]
			left++
			windows[d]--
		}
		if right-left > length {
			length = right - left
		}
	}
	return length
}

//https://leetcode.cn/problems/M1oyTv/
func minWindow1(s string, t string) string {
	need := make(map[byte]int)
	windows := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	left, right := 0, 0
	valid := 0
	start, length := 0, math.MaxInt32
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			windows[c]++
			if windows[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				length = right - left
				start = left
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if windows[d] == need[d] {
					valid--
				}
				windows[d]--
			}
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}

func Test_4(t *testing.T) {
}

//https://leetcode.cn/problems/repeated-dna-sequences/
func findRepeatedDnaSequences(s string) []string {
	seen := make(map[string]struct{})
	res := make(map[string]struct{})
	left, right := 0, 0
	for right < len(s) {
		//c := s[right]
		right++

		for right-left >= 10 {
			if right-left == 10 {
				if _, ok := seen[s[left:right]]; ok {
					res[s[left:right]] = struct{}{}
				} else {
					seen[s[left:right]] = struct{}{}
				}
			}
			//d := s[left]
			left++
		}
	}
	ans := make([]string, 0, len(res))
	for k := range res {
		ans = append(ans, k)
	}
	return ans
}
