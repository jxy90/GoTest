package middle_test

import (
	"testing"
)

func Test_romanToInt(t *testing.T) {
	//println(romanToInt("III"))
	//println(romanToInt("IV"))
	//println(romanToInt("IX"))
	//println(romanToInt("LVIII"))
	//println(romanToInt("MCMXCIV"))
	println(romanToInt("DCXXI"))
}

//var memo1 = map[string]int{
//	"I": 1,
//	"V": 5,
//	"X": 10,
//	"L": 50,
//	"C": 100,
//	"D": 500,
//	"M": 1000,
//}
//var memo = map[int]string{
//	1:    "I",
//	2:    "II",
//	3:    "III",
//	4:    "IV",
//	5:    "V",
//	9:    "IX",
//	10:   "X",
//	40:   "XL",
//	50:   "L",
//	90:   "XC",
//	100:  "C",
//	400:  "CD",
//	500:  "D",
//	900:  "CM",
//	1000: "M",
//}

var chars2 = []string{"I", "II", "III", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
var nums2 = []int{1, 2, 3, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

func romanToInt(s string) int {
	ans := 0
	index := len(chars2) - 1
	for s != "" {
		temp := chars2[index]
		n := len(temp)
		if n > len(s) {
			index--
			continue
		}
		if s[:n] == temp {
			s = s[n:]
			ans += nums2[index]
		} else {
			index--
		}
	}
	return ans
}
