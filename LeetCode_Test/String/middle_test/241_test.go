package middle_test

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_241(t *testing.T) {
	fmt.Println(diffWaysToCompute("2-1-1"))
}

var cache = map[string][]int{}

func diffWaysToCompute(expression string) []int {
	if c, ok := cache[expression]; ok {
		return c
	}
	res := make([]int, 0)
	if num, err := strconv.Atoi(expression); err == nil {
		res = append(res, num)
		cache[expression] = res
		return res
	}
	for i := range expression {
		b := expression[i]
		if b == '-' || b == '+' || b == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, leftItem := range left {
				for _, rightItem := range right {
					temp := 0
					switch b {
					case '-':
						temp = leftItem - rightItem
					case '+':
						temp = leftItem + rightItem
					case '*':
						temp = leftItem * rightItem
					}
					res = append(res, temp)
				}
			}
		}
	}
	cache[expression] = res
	return res
}
