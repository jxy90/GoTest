package easy_test

import (
	"fmt"
	"testing"
)

func Test_canWinNim(t *testing.T) {
	//nums := []int{0, 1, 0, 3, 12}
	fmt.Println(canWinNim(4))
	fmt.Println(canWinNim(8))
}

func canWinNim(n int) bool {
	return n%4 != 0
}
func canWinNim1(n int) bool {
	if n <= 3 {
		return true
	}
	f := [3]bool{true, true, true}
	for i := 4; i <= n; i++ {
		index := (i - 1) % 3
		index2 := (i - 2) % 3
		index3 := (i - 3) % 3
		f[index] = !(f[index] && f[index2] && f[index3])
	}
	return f[(n-1)%3]
}

func canWinNim0(n int) bool {
	length := n
	if length < 6 {
		length = 6
	}
	f := make([]bool, length+1)
	f[1] = true
	f[2] = true
	f[3] = true
	for i := 6; i <= n; i++ {
		f[i] = !(f[i-1] && f[i-2] && f[i-3])
	}
	return f[n]
}
