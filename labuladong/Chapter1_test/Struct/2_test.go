package Tree

import (
	"testing"
)

//type Node struct {
//	key, val  int
//	next, pre *Node
//}
//
//type DoubleNode struct {
//	head, tail *Node
//	size       int
//}
//
//func NewDoubleList() *DoubleNode {
//	head := &Node{key: 0, val: 0}
//	tail := &Node{key: 0, val: 0}
//	head.next = tail
//	tail.pre = head
//	return &DoubleNode{
//		head: head,
//		tail: tail,
//		size: 0,
//	}
//}
//
//func (d *DoubleNode) AddLast(n *Node) {
//	n.next = d.tail
//	n.pre = d.tail.pre
//	d.tail.pre.next = n
//	d.tail.pre = n
//	d.size++
//}
//
//func (d *DoubleNode) Remove(n *Node) {
//	n.pre.next = n.next
//	n.next.pre = n.pre
//	d.size--
//}
//
//func (d *DoubleNode) RemoveFirst() *Node {
//	if d.head.next == d.tail {
//		return nil
//	}
//	first := d.head.next
//	d.Remove(first)
//	return first
//}
//
////LRU缓存机制
////https://leetcode.cn/problems/lru-cache/description/
//type LRUCache struct {
//	cache map[int]*Node
//	list  *DoubleNode
//	cap   int
//}
//
//func Constructor(capacity int) LRUCache {
//	return LRUCache{
//		cache: make(map[int]*Node, capacity),
//		list:  NewDoubleList(),
//		cap:   capacity,
//	}
//}
//
//func (this *LRUCache) Get(key int) int {
//	if node, ok := this.cache[key]; ok {
//		this.makeRecently(node)
//		return node.val
//	} else {
//		return -1
//	}
//}
//
//func (this *LRUCache) makeRecently(node *Node) {
//	this.list.Remove(node)
//	this.list.AddLast(node)
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	node, ok := this.cache[key]
//	if ok {
//		//找到了，移动到tail
//		node.val = value
//		this.makeRecently(node)
//	} else {
//		//没找到，看长度。等于容量，先删除，再添加
//		if this.list.size == this.cap {
//			//删除，在此添加
//			first := this.list.RemoveFirst()
//			delete(this.cache, first.key)
//		}
//		//添加
//		node = &Node{key: key, val: value}
//		this.list.AddLast(node)
//		this.cache[key] = node
//	}
//}

func Test_2(t *testing.T) {
	//lru := Constructor(2)
	//lru.Put(1, 1)
	//lru.Put(2, 2)
	//lru.Get(1)
	//lru.Put(3, 3)
	//lru.Get(2)
	//lru.Put(4, 4)
	//lru.Get(1)
	//lru.Get(3)
	//lru.Get(4)
}
