package routine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//https://github.com/lifei6671/interview-go/blob/master/question/q011.md
//高并发下的锁与map的读写

type Ban struct {
	visitIPs map[string]time.Time
	lock     sync.Mutex
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]time.Time)}
}

func (o *Ban) visit(ip string) bool {
	o.lock.Lock()
	defer o.lock.Unlock()
	val, ok := o.visitIPs[ip]
	if ok && time.Now().Sub(val) <= time.Minute {
		return true
	}
	if !ok {
		o.visitIPs[ip] = time.Now()
	}
	if ok && time.Now().Sub(val) >= time.Minute {
		delete(o.visitIPs, ip)
	}
	return false
}

func Test_Ban(t *testing.T) {
	success := 0
	ban := NewBan()

	wg := &sync.WaitGroup{}
	wg.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func() {
				defer wg.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					success++
				}
			}()
		}

	}
	wg.Wait()
	fmt.Println("success:", success)
}
