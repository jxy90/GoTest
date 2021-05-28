package easy_test

import "testing"

func Test_exchangeBits(t *testing.T) {
	println(exchangeBits(826966453))
}

func exchangeBits(num int) int {
	even := num & 0xaaaaaaaa
	odd := num & 0x55555555
	return even>>1 | odd<<1
}
