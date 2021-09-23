package easy_test

import (
	"fmt"
	"testing"
)

func Test_isPowerOfThree(t *testing.T) {
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(0))
	fmt.Println(isPowerOfThree(9))
	fmt.Println(isPowerOfThree(45))
}

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	for n%3 == 0 {
		n = n / 3
	}
	if n == 1 {
		return true
	}
	return false
}
