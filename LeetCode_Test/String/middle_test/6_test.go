package middle_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func Test_convert(t *testing.T) {
	//fmt.Println(convert0("PAYPALISHIRING", 3))
	//fmt.Println(convert0("PAYPALISHIRING", 4))
	fmt.Println(convert0("A", 1))
	//PAYPALISHIRING
	//PA
}

func convert(s string, numRows int) string {
	n := len(s)
	diff := 2*numRows - 2
	sb := strings.Builder{}
	for i := 0; i < numRows; i++ {
		index := i
		for index < n {
			sb.WriteByte(s[index])
			index += diff
		}
	}
	return sb.String()
}
func convert0(s string, numRows int) string {
	n := len(s)
	if n == 1 || numRows >= n {
		return s
	}
	//记录保存的位置
	memo := make([][]byte, numRows)

	diff := 2*numRows - 2
	j := 0
	for i := 0; i < n; i++ {
		memo[j] = append(memo[j], s[i])
		if i%(diff) < numRows-1 {
			j++
		} else {
			j--
		}
	}

	return string(bytes.Join(memo, nil))
}
