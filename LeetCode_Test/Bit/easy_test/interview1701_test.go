package easy_test

import (
	"fmt"
	"testing"
)

func Test_add(t *testing.T) {
	fmt.Println(add(1, 2))
}

func add(a int, b int) int {
	for b != 0 {
		//tempXOR把相加不进位的位数设为1
		tempXOR := a ^ b
		//tempAnd把相加进位的位数设为1,进位1
		tempAnd := (a & b) << 1
		a = tempXOR
		b = tempAnd
	}
	return a
}
