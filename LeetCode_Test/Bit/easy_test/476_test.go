package easy_test

import (
	"fmt"
	"testing"
)

func Test_findComplement(t *testing.T) {
	fmt.Println(findComplement(2))
	fmt.Println(findComplement(1))
	fmt.Println(findComplement(5))
}

func findComplement(num int) int {
	ans := 0
	number := num
	index := 0
	for number != 0 {
		ans += (number&1 ^ 1) << index
		number >>= 1
		index++
	}
	return ans
}
