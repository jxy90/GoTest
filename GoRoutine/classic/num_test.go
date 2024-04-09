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
	wg.Add(2)
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			<-ch1
			print(i)
			ch2 <- struct{}{}
		}
		<-ch1
	}()
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			<-ch2
			print(i)
			ch1 <- struct{}{}
		}
	}()
	ch1 <- struct{}{}
	wg.Wait()
}
