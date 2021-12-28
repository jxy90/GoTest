package old

import (
	"fmt"
	"hash/crc32"
	"strconv"
	"sync"
	"testing"
)

type SyncMap []*SyncMapShared

const SHARD_COUNT = 32

type SyncMapShared struct {
	m map[string]int
	sync.RWMutex
}

func New() SyncMap {
	sm := make([]*SyncMapShared, SHARD_COUNT)
	for i := range sm {
		sm[i] = &SyncMapShared{
			m:       map[string]int{},
			RWMutex: sync.RWMutex{},
		}
	}
	return sm
}

func (sm SyncMap) getShared(key string) *SyncMapShared {
	scratch := []byte(key)
	return sm[crc32.ChecksumIEEE(scratch)%SHARD_COUNT]
}

func (sm SyncMap) Store(key string, val int) {
	shared := sm.getShared(key)
	shared.Lock()
	shared.m[key] = val
	shared.Unlock()
}
func (sm SyncMap) Load(key string) (val int, ok bool) {
	shared := sm.getShared(key)
	shared.Lock()
	val, ok = shared.m[key]
	shared.Unlock()
	return
}
func Test_ZH2(t *testing.T) {
	sm := New()
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			key := strconv.Itoa(i)
			sm.Store(key, i)
			val, _ := sm.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, val)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(sm))
}

func Test_ZH(t *testing.T) {
	fmt.Println("")
	//sync.Map{}
	sm := SyncMapShared{
		m:       map[string]int{},
		RWMutex: sync.RWMutex{},
	}
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			key := strconv.Itoa(i)
			sm.set(key, i)
			val := sm.get(key)
			fmt.Printf("k=:%v,v:=%v\n", key, val)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(sm.m))
}

func (sm *SyncMapShared) set(key string, val int) {
	defer sm.RWMutex.Unlock()
	sm.RWMutex.Lock()
	sm.m[key] = val
}

func (sm *SyncMapShared) get(key string) (ans int) {
	defer sm.RWMutex.RUnlock()
	sm.RWMutex.RLock()
	ans = sm.m[key]
	return ans
}
