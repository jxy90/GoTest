package main

import (
	"fmt"
	_ "net/http/pprof" // 引入pprof,仅使用init方法
	"sync"
)

func main() {
	var m sync.Map
	m.Store("address", map[string]string{"province": "江苏", "city": "南京"})
	v, _ := m.Load("address")
	fmt.Println(v.(map[string]string)["province"])
}
