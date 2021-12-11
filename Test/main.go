package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	runtime.GOMAXPROCS(3)
	str := "ADFasdf"
	str = strings.ToLower(str)
	fmt.Println(str)
}
