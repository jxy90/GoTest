package singleFlight

import (
	"fmt"
	"github.com/golang/groupcache/singleflight"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

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
