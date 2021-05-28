package easy_test

import "testing"

func Test_toHex(t *testing.T) {
	println(toHex(-1))
	println(toHex(26))
}
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	helper := "0123456789abcdef"
	ans := ""
	for num != 0 && len(ans) < 8 {
		temp := num & 15
		ans = string(helper[temp]) + ans
		num >>= 4
	}
	return ans
}
