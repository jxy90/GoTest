package singleFlight

import (
	"fmt"
	"github.com/golang/groupcache/singleflight"
	"log"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var group singleflight.Group

//https://developer.51cto.com/article/652064.html
//https://www.cnblogs.com/softlin/p/14133635.html
//https://blog.csdn.net/dingyu002/article/details/117958061
func Test_flightDemo(t *testing.T) {
	key := "flight"
	for i := 0; i < 5; i++ {
		log.Printf("ID: %d 请求获取缓存", i)
		go func(id int) {
			value, err := getCache(key, id)
			log.Printf("ID :%d 获取到缓存 , key: %v,value: %v,err: %v", id, key, value, err)
		}(i)
	}
	time.Sleep(20 * time.Second)
}

func getCache(key string, id int) (string, error) {
	var ret, err = group.Do(key, func() (ret interface{}, err error) {
		time.Sleep(2 * time.Second) //模拟获取缓存
		log.Printf("ID: %v 执行获取缓存", id)
		return id, nil
	})
	return strconv.Itoa(ret.(int)), err
}

/////
func TestDoDupSuppress(t *testing.T) {
	var g singleflight.Group
	c := make(chan string)
	var calls int32
	fn := func() (interface{}, error) {
		atomic.AddInt32(&calls, 1)
		return <-c, nil
	}

	const n = 10
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() { // n个协程同时调用了g.Do，fn中的逻辑只会被一个协程执行
			v, err := g.Do("key", fn)
			if err != nil {
				t.Errorf("Do error: %v \r\n", err)
			}
			if v.(string) != "bar" {
				t.Errorf("got %q; want %q \n", v, "bar")
			}
			wg.Done()
		}()
	}
	time.Sleep(100 * time.Millisecond) // let goroutines above block
	c <- "bar"
	wg.Wait()
	got := atomic.LoadInt32(&calls)
	if got != 1 {
		t.Errorf("number of calls = %d; want 1 \n", got)
	}
	fmt.Println(got)
}
