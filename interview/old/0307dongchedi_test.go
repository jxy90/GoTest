package old

import (
	"fmt"
	"sync"
	"testing"
)

func Test_dongchedi(t *testing.T) {
	//fmt.Println("123")
	c1 := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 1; i <= 100; i++ {
			c1 <- 1
			if i%2 == 1 {
				fmt.Println(i)
			}
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i <= 100; i++ {
			<-c1
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
		wg.Done()
	}()
	//c1 <- 1
	wg.Wait()
}
