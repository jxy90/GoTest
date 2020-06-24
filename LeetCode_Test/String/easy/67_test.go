package easy_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"strconv"
	"testing"
)

func addBinary(a string, b string) string {
	ans := ""
	lena := len(a)
	lenb := len(b)
	n := CommonUtil.Max(lena, lenb)
	carry := 0
	for i := 0; i < n; i++ {
		if i < lena {
			carry += int(a[lena-i-1] - '0')
		}
		if i < lenb {
			carry += int(b[lenb-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry = carry / 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

func Test_addBinary(t *testing.T) {
	a, b := "11", "1"
	data := addBinary(a, b)
	println(data)
	a, b = "1010", "1011"
	data = addBinary(a, b)
	println(data)
}
