package main

import (
	"fmt"
)

// Print 泛型
func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
}

func main() {
	Print([]string{"你好, ", "脑子进了煎鱼\n"})
	Print([]int64{1, 2, 3})
}
