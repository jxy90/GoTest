package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // 引入pprof,仅使用init方法
	"time"
)

func main() {

	// 生产环境应仅在本地监听pprof
	go func() {
		ip := "127.0.0.1:9527"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Println("开启pprof失败", ip, err)
		}
	}()

	// 业务代码运行
	outCh := make(chan int)
	// 死代码，永不读取
	go func() {
		if false {
			<-outCh
		}
		select {}
	}()

	// 每秒起10个goroutine，goroutine会阻塞，不释放内存
	tick := time.Tick(time.Second / 10)
	i := 0
	for range tick {
		i++
		fmt.Println(i)
		alloc1(outCh) // 不停的有goruntine因为outCh堵塞，无法释放
	}
}

// 一个外层函数
func alloc1(outCh chan<- int) {
	go alloc2(outCh)
}

// 一个内层函数
func alloc2(outCh chan<- int) {
	func() {
		defer fmt.Println("alloc-fm exit")
		// 分配内存，假用一下
		buf := make([]byte, 1024*1024*10)
		_ = len(buf)
		fmt.Println("alloc done")

		outCh <- 0 // 54行
	}()
}
