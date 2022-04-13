package Struct

import (
	"golang.org/x/exp/rand"
	"testing"
)

func Test_RandomizedSet(t *testing.T) {

}

type RandomizedSet struct {
	m map[int]struct{}
	s []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m: map[int]struct{}{},
		s: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.m[val] = struct{}{}
	this.s = append(this.s, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}
	delete(this.m, val)
	for i := range this.s {
		if this.s[i] == val {
			if i == len(this.s) {
				this.s = this.s[:i]
			} else {
				this.s = append(this.s[:i], this.s[i+1:]...)
			}
			break
		}
	}
	return true
}

func (this *RandomizedSet) GetRandom() int {
	rand := rand.Intn(len(this.s))
	return this.s[rand]
}
