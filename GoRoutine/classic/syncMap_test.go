package classic

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

//实现阻塞读且并发安全的map
//https://github.com/lifei6671/interview-go/blob/master/question/q010.md
func Test_Map(t *testing.T) {

	mapval := Map{
		Cache:  make(map[string]*entry),
		Locker: &sync.RWMutex{},
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			val := mapval.Read(fmt.Sprint(i), time.Second*6)
			log.Println("读取值为->", val)
		}(i)
	}

	time.Sleep(time.Second * 3)
	for i := 0; i < 10; i++ {
		go func(val int) {
			mapval.Write(fmt.Sprint(val), val)
		}(i)
	}

	time.Sleep(time.Second * 30)
}

type sp interface {
	Write(key string, val interface{})                  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Read(key string, timeout time.Duration) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type Map struct {
	Cache  map[string]*entry
	Locker *sync.RWMutex
}
type entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}

func (m *Map) Write(key string, val interface{}) {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	item, ok := m.Cache[key]
	if !ok {
		m.Cache[key] = &entry{
			value:   val,
			isExist: true,
		}
		return
	}
	item.value = val
	if !item.isExist {
		if item.ch != nil {
			close(item.ch)
			item.ch = nil
		}
	}
	return
}

func (m *Map) Read(key string, timeout time.Duration) interface{} {
	m.Locker.RLock()
	if e, ok := m.Cache[key]; ok && e.isExist {
		m.Locker.RUnlock()
		return e.value
	} else if !ok {
		m.Locker.RUnlock()
		m.Locker.Lock()
		e = &entry{ch: make(chan struct{}), isExist: false}
		m.Cache[key] = e
		m.Locker.Unlock()
		log.Println("协程阻塞 -> ", key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			log.Println("协程超时 -> ", key)
			return nil
		}
	} else {
		m.Locker.RUnlock()
		log.Println("协程阻塞 -> ", key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			log.Println("协程超时 -> ", key)
			return nil
		}
	}
}
