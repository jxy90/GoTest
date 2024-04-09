package classic

import (
	"sync"
	"testing"
)

//循环打印数字
//12345678910
func Test10(t *testing.T) {
	print10()
}

func print10() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ch1 := make(chan int)
	ch2 := make(chan int)
	defer func() {
		close(ch1)
		close(ch2)
		ch1 = nil
		ch2 = nil
	}()
	go func() {
		for {
			select {
			case val := <-ch1:
				if val > 10 {
					wg.Done()
					break
				}
				print(val)
				ch2 <- val + 1
			}
		}
	}()
	go func() {
		for {
			select {
			case val := <-ch2:

				print(val)
				ch1 <- val + 1
			}
		}
	}()

	ch1 <- 1
	wg.Wait()
}
