package routine

import (
	"fmt"
	"testing"
	"time"
)

func Test2(t *testing.T) {
	ch0 := make(chan struct{})
	ch1 := make(chan struct{})
	i := 0
	for {
		if i%2 == 0 {
			go func(val int) {
				<-ch0
				fmt.Println(val)
				ch1 <- struct{}{}
			}(i)
		} else {
			go func(val int) {
				<-ch1
				fmt.Println(val)
				ch0 <- struct{}{}
			}(i)
		}
		fmt.Println(i)
		i++
		time.Sleep(time.Second)
	}
	select {}
}

//func Test2(t *testing.T) {
//	//两个goroutine，分别打印奇数和偶数
//	wg := &sync.WaitGroup{}
//	ch0 := make(chan struct{})
//	ch1 := make(chan struct{})
//	wg.Add(2)
//	go func() {
//		defer wg.Done()
//		for i := 0; i < 10; i++ {
//			<-ch0
//			fmt.Println("0")
//			ch1 <- struct{}{}
//		}
//		close(ch0)
//	}()
//
//	go func() {
//		defer wg.Done()
//		for i := 0; i < 10; i++ {
//			<-ch1
//			fmt.Println("1")
//			if i != 9 {
//				ch0 <- struct{}{}
//			}
//		}
//		close(ch1)
//	}()
//
//	ch0 <- struct{}{}
//	wg.Wait()
//}
