package classic

import (
	"fmt"
	"sync"
	"testing"
)

var ch = make(chan struct{})
var wg sync.WaitGroup

func go1() {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
		ch <- struct{}{} //不能与上一行交换位置
		<-ch
	}
}
func go2() {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<-ch
		fmt.Println(i)
		ch <- struct{}{}
	}
}
func Test2(t *testing.T) {
	ch := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			ch <- struct{}{}
			//奇数
			if i%2 == 1 {
				fmt.Println("线程1打印:", i)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			<-ch
			//偶数
			if i%2 == 0 {
				fmt.Println("线程2打印:", i)
			}
		}
		close(ch)
	}()
	wg.Wait()
	//select {}
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
