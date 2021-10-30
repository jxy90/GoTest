package main

import (
	"fmt"
	"testing"
)

func Test_Continue(t *testing.T) {
	// 不使用标记
	fmt.Println("---- continue ---- ")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			continue
		}
	}

	// 使用标记
	fmt.Println("---- continue label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i11: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i21: %d\n", i2)
			continue re
			fmt.Printf("i22: %d\n", i2)
		}
		fmt.Printf("i12: %d\n", i)
	}
	// break
	fmt.Println("---- continue label ----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i11: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i21: %d\n", i2)
			break
			fmt.Printf("i22: %d\n", i2)
		}
		fmt.Printf("i12: %d\n", i)
	}
}
