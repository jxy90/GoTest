package middle_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_countAndSay(t *testing.T) {
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
	//fmt.Println(countAndSayHelper("11"))
	//fmt.Println(countAndSayHelper("21"))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	ans := "11"
	for i := 3; i <= n; i++ {
		ans = countAndSayHelper(ans)
	}
	return ans
}

func countAndSayHelper(str string) string {
	ans := strings.Builder{}
	n := len(str)
	j := 0
	for i := 0; i <= n; i++ {
		if i == n || str[j] != str[i] {
			count := i - j
			ans.WriteString(strconv.Itoa(count))
			ans.WriteByte(str[j])
			j = i
		}
	}
	return ans.String()
}
