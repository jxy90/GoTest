package classic

import "testing"

func Test_LFU(t *testing.T) {
	//lfu := Constructor(2)
	//lfu.Put(1, 1)
	//lfu.Put(2, 2)
	//println(lfu.Get(1))
	//lfu.Put(3, 3)
	//println(lfu.Get(2))
	//println(lfu.Get(3))
	//lfu.Put(4, 4)
	//println(lfu.Get(1))
	//println(lfu.Get(3))
	//println(lfu.Get(4))

	//lfu := Constructor(1)
	//lfu.Put(0, 0)
	//println(lfu.Get(0))

	//lfu := Constructor(3)
	//lfu.Put(2, 2)
	//lfu.Put(1, 1)
	//println(lfu.Get(2))
	//println(lfu.Get(1))
	//println(lfu.Get(2))
	//lfu.Put(3, 3)
	//lfu.Put(4, 4)
	//println(lfu.Get(3))
	//println(lfu.Get(2))
	//println(lfu.Get(1))
	//println(lfu.Get(4))
}

//type Node struct {
//	pre, next       *Node
//	key, val, level int
//}
//
//type DoubleList struct {
//	head, tail *Node
//}
//
//func newDoubleList() DoubleList {
//	head := &Node{}
//	tail := &Node{}
//	head.next = tail
//	tail.pre = head
//	return DoubleList{
//		head: head,
//		tail: tail,
//	}
//}
//func (list *DoubleList) add(node *Node) {
//	node.next = list.head.next
//	node.pre = list.head
//	list.head.next.pre = node
//	list.head.next = node
//}
//
//func (list DoubleList) remove(node *Node) {
//	node.pre.next = node.next
//	node.next.pre = node.pre
//}
//
//type LFUCache struct {
//	cache               map[int]*Node
//	lists               map[int]*DoubleList
//	size, cap, minLevel int
//}
//
//func Constructor(capacity int) LFUCache {
//	return LFUCache{
//		cache: make(map[int]*Node),
//		lists: make(map[int]*DoubleList),
//		size:  0,
//		cap:   capacity,
//	}
//}
//
//func (this *LFUCache) Get(key int) int {
//	node, ok := this.cache[key]
//	if !ok {
//		return -1
//	}
//	this.remove(node)
//	node.level++
//	this.add(node)
//	return node.val
//}
//
//func (this *LFUCache) Put(key int, value int) {
//	node, ok := this.cache[key]
//	if ok {
//		node.val = value
//		this.Get(key)
//		return
//	}
//	node = &Node{
//		key:   key,
//		val:   value,
//		level: 0,
//	}
//	this.add(node)
//}
//
//func (this *LFUCache) add(node *Node) {
//	list, ok := this.lists[node.level]
//	if !ok {
//		temp := newDoubleList()
//		list = &temp
//		this.lists[node.level] = list
//	}
//	if this.cap == this.size {
//		this.removeLast()
//	}
//	list.add(node)
//	this.cache[node.key] = node
//	this.size++
//	if node.level == 0 {
//		this.minLevel = 0
//	}
//}
//
//func (this *LFUCache) remove(node *Node) {
//	list, _ := this.lists[node.level]
//	list.remove(node)
//	delete(this.cache, node.key)
//	this.size--
//	for list != nil && list.head.next == list.tail && this.minLevel == node.level {
//		this.minLevel++
//		list, _ = this.lists[this.minLevel]
//	}
//}
//
//func (this *LFUCache) removeLast() *Node {
//	list, _ := this.lists[this.minLevel]
//	node := list.tail.pre
//	this.remove(node)
//	return node
//}
