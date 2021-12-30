package main

import (
	"fmt"
	"unsafe"
)

// 空结构体
var Exists = struct{}{}

// Set is the main interface
type Set struct {
	// struct为结构体类型的变量
	m map[interface{}]struct{}
}

// 定义非空结构体
type S struct {
	a uint16
	b uint32
}

func main() {
	var s S
	fmt.Println(unsafe.Sizeof(s)) // prints 8, not 6
	var s2 struct{}
	fmt.Println(unsafe.Sizeof(s2)) // prints 0
}

func New(items ...interface{}) *Set {
	// 获取Set的地址
	s := &Set{}
	// 声明map类型的数据结构
	s.m = make(map[interface{}]struct{})
	s.Add(items...)
	return s
}

func (s *Set) Add(items ...interface{}) error {
	for _, item := range items {
		s.m[item] = Exists
	}
	return nil
}
func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}
func (s *Set) Size() int {
	return len(s.m)
}
func (s *Set) Clear() {
	s.m = make(map[interface{}]struct{})
}
func (s *Set) Equal(other *Set) bool {
	// 如果两者Size不相等，就不用比较了
	if s.Size() != other.Size() {
		return false
	}

	// 迭代查询遍历
	for key := range s.m {
		// 只要有一个不存在就返回false
		if !other.Contains(key) {
			return false
		}
	}
	return true
}
func (s *Set) IsSubset(other *Set) bool {
	// s的size长于other，不用说了
	if s.Size() > other.Size() {
		return false
	}
	// 迭代遍历
	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}
