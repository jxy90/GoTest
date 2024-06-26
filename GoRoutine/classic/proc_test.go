package classic

import (
	"fmt"
	"testing"
	"time"
)

//https://github.com/lifei6671/interview-go/blob/master/question/q012.md
//1. 写出以下逻辑，要求每秒钟调用一次proc并保证程序不退出?

func Test_proc(t *testing.T) {
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出
		t := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-t.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					proc()
				}()
			}
		}
	}()

	select {}
}

func proc() {
	panic("ok")
}
