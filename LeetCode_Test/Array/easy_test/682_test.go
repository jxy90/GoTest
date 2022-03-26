package easy_test

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_682(t *testing.T) {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
}

func calPoints(ops []string) int {
	n := len(ops)
	res := make([]int, 0, n)
	for _, op := range ops {
		i := 0
		switch op {
		case "+":
			i = res[len(res)-1] + res[len(res)-2]
		case "D":
			i = 2 * res[len(res)-1]
		case "C":
			res = res[:len(res)-1]
		default:
			i, _ = strconv.Atoi(op)
		}
		if op != "C" {
			res = append(res, i)
		}
	}
	ans := 0
	//fmt.Println(res)
	for _, re := range res {
		ans += re
	}
	return ans
}
