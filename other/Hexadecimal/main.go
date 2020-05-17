package main

import (
	"fmt"
)

func main() {
	i := 155
	h := fmt.Sprintf("%x", i)
	fmt.Printf("Hex conv of '%d' is '%s'\n", i, h)
	h = fmt.Sprintf("%X", i)
	fmt.Printf("HEX conv of '%d' is '%s'\n", h, h)

	fmt.Println(toHex(i))
}

func toHex(num int) string {
	const (
		hexStr = "0123456789abcdef"
		mask   = 0xf
	)
	var res = make([]byte, 8)
	var nonZeroIndex = 7
	for i := 7; i >= 0; i-- {
		Val := num & mask
		res[i] = hexStr[Val]
		if Val > 0 {
			nonZeroIndex = i
		}
		num >>= 4
	}
	return string(res[nonZeroIndex:])
}
