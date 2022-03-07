package old

import (
	"fmt"
	"sync"
	"testing"
)

func Test_0(t *testing.T) {
	//fmt.Println("123")
	c1 := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i += 2 {
			c1 <- 1
			fmt.Println(i)
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i < 10; i += 2 {
			<-c1
			fmt.Println(i)
		}
		wg.Done()
	}()
	//c1 <- 1
	wg.Wait()
}
