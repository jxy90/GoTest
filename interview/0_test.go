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
		for i := 0; i < 10; i++ {

			c1 <- 1

			if i%2 == 0 {
				fmt.Print(i)
			}
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10; i++ {

			<-c1
			if i%2 == 1 {
				fmt.Print(i)
			}
		}
		wg.Done()
	}()
	//c1 <- 1
	wg.Wait()
}
