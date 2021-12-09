package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(3)
	fmt.Println(runtime.NumGoroutine())
}
