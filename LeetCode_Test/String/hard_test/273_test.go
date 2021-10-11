package hard_test

import (
	"fmt"
	"strings"
	"testing"
)

func Test_numberToWords(t *testing.T) {
	fmt.Println(numberToWords(123))
}

var num2str_small = []string{
	"Zero",
	"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
	"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen",
}
var num2str_medium = []string{
	"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety",
}

var num2str_large = []string{
	"Billion", "Million", "Thousand", "",
}

func numberToWords(num int) string {
	if num == 0 {
		return num2str_small[0]
	}
	sb := strings.Builder{}
	for i, j := 1000000000, 0; i > 0; i, j = i/1000, j+1 {
		if num < i {
			continue
		}
		sb.WriteString(numToStr(num/i) + num2str_large[j] + " ")
		num %= i
	}
	return strings.TrimSpace(sb.String())

}

func numToStr(x int) (ans string) {
	if x >= 100 {
		ans += num2str_small[x/100] + " Hundred "
		x %= 100
	}
	if x >= 20 {
		ans += num2str_medium[x/10] + " "
		x %= 10
	}
	if x != 0 {
		ans += num2str_small[x] + " "
	}
	return
}
