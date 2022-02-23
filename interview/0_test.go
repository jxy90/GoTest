package old

import (
	"fmt"
	"testing"
)

var ch chan int

func Test_0(t *testing.T) {
	ch = make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	for i := range ch {
		fmt.Println(i)
	}
	//close(ch)
}

func Test_1(t *testing.T) {
	ch := make(chan int) //创建一个无缓存channel

	//新建一个goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //往通道写数据
		}
		//不需要再写数据时，关闭channel
		close(ch)
		//ch <- 666 //关闭channel后无法再发送数据

	}() //别忘了()

	for num := range ch {
		fmt.Println("num = ", num)
	}
}
