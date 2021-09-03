package easy_test

import (
	"bytes"
	"strings"
	"testing"
)

func Test_reverseVowels(t *testing.T) {
	fmt.Println(reverseVowels("hello"))
}

func reverseVowels(s string) string {
	str := "aeiouAEIOU"
	sBytes := []byte(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right {
			if strings.ContainsRune(str, rune(s[left])) {
				break
			}
			left++
		}
		for left < right {
			if strings.ContainsRune(str, rune(s[right])) {
				break
			}
			right--
		}
		if left < right {
			sBytes[left], sBytes[right] = sBytes[right], sBytes[left]
			left++
			right--
		}
	}
	return string(sBytes)
}

func reverseVowels0(s string) string {
	str := "aeiouAEIOU"
	nums := make([]int, 0)
	for i, v := range s {
		if bytes.ContainsRune([]byte(str), v) {
			nums = append(nums, i)
		}
	}
	left, right := 0, len(nums)-1
	sBytes := []byte(s)
	for left < right {
		sBytes[nums[left]], sBytes[nums[right]] = sBytes[nums[right]], sBytes[nums[left]]
		left++
		right--
	}
	return string(sBytes)
}
