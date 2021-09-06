package easy_test

import (
	"fmt"
	"testing"
)

func Test_convertInteger(t *testing.T) {
	//2
	fmt.Println(convertInteger(29, 15))
	//14
	fmt.Println(convertInteger(826966453, -729934991))
}

func convertInteger(A int, B int) int {
	temp := int32(A ^ B)
	count := 0
	for temp != 0 {
		temp = temp & (temp - 1)
		count++
	}
	return count
}
