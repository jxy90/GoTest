package easy_test

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_504(t *testing.T) {
	fmt.Println(convertToBase7(100))
	fmt.Println(convertToBase7(-7))
	fmt.Println(convertToBase7(-8))
}

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	c := num
	flag := true
	if num < 0 {
		flag = false
		c = -c
	}
	ints := make([]int, 0)
	for c != 0 {
		ints = append(ints, c%7)
		c /= 7
	}
	s := ""
	if !flag {
		s = "-"
	}
	for i := len(ints) - 1; i >= 0; i-- {
		s += strconv.Itoa(ints[i])
	}
	return s
}
