package cmap

import (
	"fmt"
	"strconv"
	"sync"
)

var syncMap sync.Map

func main() {
	wait := sync.WaitGroup{}
	for i := 0; i < 3000; i++ {
		wait.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			syncMap.Store(key, n)
			value, _ := syncMap.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wait.Done()
		}(i)

	}
	wait.Wait()
}
