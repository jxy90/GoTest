package hard_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func Test_564(t *testing.T) {
	fmt.Println(nearestPalindromic("123"))
	fmt.Println(nearestPalindromic("1"))
	fmt.Println(nearestPalindromic("1234"))
	fmt.Println(nearestPalindromic("999"))
}

func nearestPalindromic(n string) string {
	m := len(n)
	options := []int{int(math.Pow10(m-1) - 1), int(math.Pow10(m) + 1)}
	half := (m + 1) / 2
	halfVal, _ := strconv.Atoi(n[:half])
	for _, v := range []int{halfVal + 1, halfVal, halfVal - 1} {
		temp := v
		if m%2 == 1 {
			temp /= 10
		}
		for temp > 0 {
			v = v*10 + temp%10
			temp /= 10
		}
		options = append(options, v)
	}
	fmt.Println(options)
	ans := -1
	origin, _ := strconv.Atoi(n)
	for _, option := range options {
		if option == origin {
			continue
		}
		if ans == -1 || abs(ans-origin) > abs(option-origin) || (abs(ans-origin) == abs(option-origin) && option < ans) {
			ans = option
		}
	}
	return strconv.Itoa(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
