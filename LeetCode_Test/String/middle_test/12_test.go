package middle_test

import (
	"strings"
	"testing"
)

func Test_intToRoman(t *testing.T) {
	println(intToRoman(3))
	println(intToRoman(4))
	println(intToRoman(9))
	println(intToRoman(58))
	println(intToRoman(1994))
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

var chars = []string{"I", "II", "III", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
var nums = []int{1, 2, 3, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

func intToRoman(num int) string {
	ans := strings.Builder{}
	index := len(chars) - 1
	for num != 0 {
		if num-nums[index] >= 0 {
			num -= nums[index]
			ans.WriteString(chars[index])
		} else {
			index--
		}
	}
	return ans.String()
}
