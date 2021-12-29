package routine

import (
	"log"
	"sync"
	"testing"
	"time"
)

//https://github.com/lifei6671/interview-go/blob/master/question/q010.md
//实现阻塞读且并发安全的map

type sp interface {
	Out(key string, val interface{})                  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type entry struct {
	ch      chan struct{}
	val     interface{}
	isExist bool
}
type Map struct {
	entries map[string]*entry
	lock    *sync.RWMutex
}

func (m *Map) Out(key string, val interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	item, ok := m.entries[key]
	if !ok {
		m.entries[key] = &entry{
			val:     val,
			isExist: true,
		}
	}
	item.val = val
	if !item.isExist && item.ch != nil {
		close(item.ch)
		item.ch = nil
	}
}

func (m *Map) Rd(key string, timeout time.Duration) interface{} {
	m.lock.RLock()
	//defer
	item, ok := m.entries[key]
	m.lock.RUnlock()
	if ok && item.isExist {
		return item.val
	} else if !ok {
		m.lock.Lock()
		item = &entry{
			ch:      make(chan struct{}),
			isExist: false,
		}
		m.entries[key] = item
		m.lock.Unlock()
		log.Println("协程阻塞 -> ", key)
		select {
		case <-item.ch:
			return item.val
		case <-time.After(timeout):
			return nil
		}
	} else {
		log.Println("协程阻塞 -> ", key)
		select {
		case <-item.ch:
			return item.val
		case <-time.After(timeout):
			return nil
		}
	}
}

func Test_map1(t *testing.T) {
	mapval := Map{
		entries: make(map[string]*entry),
		lock:    &sync.RWMutex{},
	}

	for i := 0; i < 10; i++ {
		go func() {
			val := mapval.Rd("key", time.Second*6)
			log.Println("读取值为->", val)
		}()
	}

	time.Sleep(time.Second * 3)
	for i := 0; i < 10; i++ {
		go func(val int) {
			mapval.Out("key", val)
		}(i)
	}

	time.Sleep(time.Second * 5)
}
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
