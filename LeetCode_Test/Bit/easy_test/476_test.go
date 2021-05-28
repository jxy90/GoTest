package easy_test

import "testing"

func Test_findComplement(t *testing.T) {
	println(findComplement(2))
	println(findComplement(1))
	println(findComplement(5))
}

func findComplement(num int) int {
	ans := 0
	index := 0
	for num != 0 {
		ans = ans + (num&1^1)<<index
		num >>= 1
		index++
	}
	return ans
}
