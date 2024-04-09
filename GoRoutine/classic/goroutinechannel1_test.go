package classic

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

//https://github.com/lifei6671/interview-go/blob/master/question/q009.md
//在 golang 协程和channel配合使用
//写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。

func Test1(t *testing.T) {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- rand.Intn(5)
		}
		close(ch)
	}()
	go func() {
		defer wg.Done()
		//for i := 0; i < 5; i++ {
		//	fmt.Println(<-ch)
		//}
		for i := range ch {
			fmt.Println(i)
		}
	}()
	wg.Wait()
}
