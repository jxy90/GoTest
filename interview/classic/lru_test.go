package classic

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

//https://leetcode.cn/problems/lru-cache/description/?envType=list&envId=xcwvuni
func Test(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	assert.Equal(t, lru.Get(1), 1)
	lru.Put(3, 3)
	assert.Equal(t, (lru.Get(2)), -1)
	lru.Put(4, 4)
	assert.Equal(t, (lru.Get(1)), -1)
	assert.Equal(t, (lru.Get(3)), 3)
	assert.Equal(t, (lru.Get(4)), 4)
}

type Node struct {
	pre, next *Node
	key, val  int
}

type DoubleList struct {
	head, tail *Node
}

func newDoubleList() DoubleList {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.next = head
	return DoubleList{head, tail}
}

func (list DoubleList) add(node *Node) {
	node.pre = list.head
	node.next = list.head.next
	list.head.next.pre = node
	list.head.next = node
}
func (list DoubleList) remove(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}
func (list DoubleList) removeTail() *Node {
	node := list.tail.pre
	list.remove(node)
	return node
}

type LRUCache struct {
	cache     map[int]*Node
	list      *DoubleList
	size, cap int
}

func Constructor(capacity int) LRUCache {
	list := newDoubleList()
	return LRUCache{
		cache: make(map[int]*Node),
		list:  &list,
		size:  0,
		cap:   capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.list.remove(node)
	this.list.add(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		node.val = value
		this.Get(key)
		return
	}
	node = &Node{
		key: key,
		val: value,
	}
	if this.size == this.cap {
		tail := this.list.removeTail()
		delete(this.cache, tail.key)
		this.size--
	}
	this.list.add(node)
	this.cache[key] = node
	this.size++
}
