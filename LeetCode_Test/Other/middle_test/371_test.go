package middle_test

import (
	"fmt"
	"testing"
)

func Test_getSum(t *testing.T) {
	fmt.Println(getSum(1, 2))
	fmt.Println(getSum(3, 2))
	fmt.Println(getSum(-3, -2))
}

func getSum(a int, b int) int {
	for b != 0 {
		temp := a & b << 1
		a ^= b
		b = temp
	}
	return a
}
