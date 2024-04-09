package _024

import (
	"github.com/go-playground/assert/v2"
	"sync"
	"testing"
)

func TestBitMap(t *testing.T) {
	bm := NewBitMap(4)
	bm.Set(1)
	assert.Equal(t, bm.Check(1), true)
	bm.Unset(1)
	assert.Equal(t, bm.Check(1), false)
}

type BitMap struct {
	locker *sync.RWMutex
	bits   []byte
}

func NewBitMap(max_val uint) *BitMap {
	size := (max_val + 7) / 8
	return &BitMap{
		locker: &sync.RWMutex{},
		bits:   make([]byte, size),
	}
}

func (bm *BitMap) Set(num uint) {
	if num >= uint(len(bm.bits)*8) {
		panic("number out of bounds")
	}
	bm.locker.Lock()
	defer bm.locker.Unlock()
	index := num / 8
	bitOffset := index % 8
	bm.bits[index] |= 1 << bitOffset
}
func (bm *BitMap) Unset(num uint) {
	if num >= uint(len(bm.bits)*8) {
		panic("number out of bounds")
	}
	bm.locker.Lock()
	defer bm.locker.Unlock()
	index := num / 8
	bitOffset := index % 8
	bm.bits[index] &= ^(1 << bitOffset)
}
func (bm *BitMap) Check(num uint) bool {
	bm.locker.RLock()
	defer bm.locker.RUnlock()
	index := num / 8
	bitOffset := index % 8
	return bm.bits[index]&(1<<bitOffset) > 0
}
