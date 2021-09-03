package easy_test

import "testing"

func Test_findComplement(t *testing.T) {
	fmt.Println(findComplement(2))
	fmt.Println(findComplement(1))
	fmt.Println(findComplement(5))
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
