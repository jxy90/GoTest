package hard_test

import (
	"fmt"
	"strings"
	"testing"
)

func Test_isNumber(t *testing.T) {
	tests := []string{"2e0", "2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"}
	tests = []string{"e.7e5", "4..", ".", "abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"}
	for i := range tests {
		//temp := "2e0"
		//fmt.Println(fmt.Sprintf("%-15v string is : %-5v", temp, isNumber(temp)))
		fmt.Println(fmt.Sprintf("%-15v string is : %-5v", tests[i], isNumber(tests[i])))
	}
}

func isNumber(s string) bool {
	//根据e分割左边是整数/浮点数,最右只能是整数
	strs := make([]string, 0)
	se := strings.Split(s, "e")
	counte := len(se) - 1
	for i := range se {
		SE := strings.Split(se[i], "E")
		counte += len(SE) - 1
		for j := range SE {
			if SE[j] != "" {
				strs = append(strs, SE[j])
			}
		}
	}
	n := len(strs)
	if counte > 1 {
		return false
	}
	if n == 2 {
		return isValid(strs[0], false) && isValid(strs[1], true)
	}
	return isValid(s, false)
}

func isValid(s string, mustInt bool) bool {
	n := len(s)
	i := 0
	if isSignValid(s[0]) {
		i++
	}
	hasPoint, hasNum := false, false
	for i < n {
		b := s[i]
		if isPointValid(b) {
			if mustInt || hasPoint {
				return false
			}
			hasPoint = true
		} else if isNumberValid(b) {
			hasNum = true
		} else {
			return false
		}
		i++
	}
	return hasNum
}

//func isNumber0(s string) bool {
//	//没通过该
//	n := len(s)
//	if n == 1 {
//		return isNumberValid(s[0])
//	}
//	hasNumber := false
//	hasE := false
//	for i := 0; i < n; i++ {
//		b := s[i]
//		if isEValid(b) && i != 0 && i+1 < n {
//			//e开头
//			hasE = true
//			b1 := s[i+1]
//			if isNumberValid(b1) {
//				continue
//			} else if isSignValid(b1) && i+2 < n && isNumberValid(s[i+2]) {
//				continue
//			} else {
//				return false
//			}
//		} else if isSignValid(b) && i+1 < n && (isNumberValid(s[i+1]) || isPointValid(s[i+1])) {
//			//符号开头
//			continue
//		} else if isPointValid(b) && hasE == false {
//			//.开头
//			if i-1 >= 0 && isNumberValid(s[i-1]) {
//				if i+1 >= n {
//					return true
//				}
//				continue
//			} else if i+1 < n && isNumberValid(s[i+1]) {
//				continue
//			} else {
//				return false
//			}
//		} else if isNumberValid(b) {
//			//数字开头
//			hasNumber = true
//			continue
//		} else {
//			return false
//		}
//	}
//	if !hasNumber {
//		return false
//	}
//	return true
//}

func isNumberValid(b byte) bool {
	if '0' <= b && b <= '9' {
		return true
	}
	return false
}
func isEValid(b byte) bool {
	if 'E' == b || 'e' <= b {
		return true
	}
	return false
}
func isSignValid(b byte) bool {
	if '+' == b || '-' == b {
		return true
	}
	return false
}
func isPointValid(b byte) bool {
	return '.' == b
}
