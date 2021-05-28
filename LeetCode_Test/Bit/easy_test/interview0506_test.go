package easy_test

import "testing"

func Test_convertInteger(t *testing.T) {
	//2
	println(convertInteger(29, 15))
	//14
	println(convertInteger(826966453, -729934991))
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
