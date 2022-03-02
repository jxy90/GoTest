package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 0)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)
}
