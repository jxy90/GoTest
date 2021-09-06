package easy_test

import (
	"fmt"
	"testing"
)

func Test_exchangeBits(t *testing.T) {
	fmt.Println(exchangeBits(826966453))
}

func exchangeBits(num int) int {
	even := num & 0xaaaaaaaa
	odd := num & 0x55555555
	return even>>1 | odd<<1
}
