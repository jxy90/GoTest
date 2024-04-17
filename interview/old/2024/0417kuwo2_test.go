package _024

import (
	"github.com/go-playground/assert/v2"
	"log"
	"testing"
)

func Test_LRU(t *testing.T) {
	lru := newLRU(2)
	lru.Put("1", "1")
	lru.Put("2", "2")
	assert.Equal(t, lru.Get("1"), "1")
	assert.Equal(t, lru.Get("2"), "2")
	lru.Put("3", "3")
	assert.Equal(t, lru.Get("3"), "3")
	assert.Equal(t, lru.Get("1"), "")
	assert.Equal(t, lru.Get("2"), "2")
}

//lru
//put, get

type Node struct {
	pre, next *Node
	key, val  string
}

type NodeList struct {
	head, tail *Node
}

func newNodeList() *NodeList {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.pre = head
	return &NodeList{
		head: head,
		tail: tail,
	}
}

func (list *NodeList) Add(node *Node) {
	node.next = list.head.next
	node.pre = list.head
	list.head.next.pre = node
	list.head.next = node
}
func (list *NodeList) Remove(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}
func (list *NodeList) RemoveLast() *Node {
	node := list.tail.pre
	list.Remove(node)
	return node
}

type LRU struct {
	cache     map[string]*Node
	list      *NodeList
	size, cap int
}

func newLRU(cap int) LRU {
	return LRU{
		cache: make(map[string]*Node),
		list:  newNodeList(),
		cap:   cap,
	}
}

func (l *LRU) Get(key string) string {
	node, ok := l.cache[key]
	if !ok {
		log.Printf("key: %s not exist!!!", key)
		return ""
	}
	l.list.Remove(node)
	l.list.Add(node)
	return node.val
}

func (l *LRU) Put(key, val string) {
	node, ok := l.cache[key]
	if ok {
		node.val = val
		l.Get(node.key)
		return
	}
	if l.cap == l.size {
		last := l.list.RemoveLast()
		delete(l.cache, last.key)
		l.size--
	}
	node = &Node{
		key: key,
		val: val,
	}
	l.list.Add(node)
	l.cache[node.key] = node
	l.size++
}
