package Tree

import (
	"testing"
)

//type DoubleNode struct {
//	key, val, level int
//	next, pre       *DoubleNode
//}
//
//type DoubleNodeList struct {
//	head, tail *DoubleNode
//}
//
//type LFUCache struct {
//	cap, size, minLevel int
//	levelDoubleNodeList map[int]*DoubleNodeList
//	cache               map[int]*DoubleNode
//}
//
//func Constructor(capacity int) LFUCache {
//	return LFUCache{
//		cap:                 capacity,
//		minLevel:            1,
//		levelDoubleNodeList: make(map[int]*DoubleNodeList),
//		cache:               make(map[int]*DoubleNode),
//	}
//}
//
//func (this *LFUCache) addNewNode(key, val int) {
//	node := &DoubleNode{key: key, val: val, level: 1}
//	this.addNode(node)
//	this.minLevel = 1
//}
//
//func (this *LFUCache) addNode(node *DoubleNode) {
//	nodeLevel := node.level
//	if this.levelDoubleNodeList[nodeLevel] == nil {
//		this.levelDoubleNodeList[nodeLevel] = this.initDoubleList()
//	}
//	levelNodeList := this.levelDoubleNodeList[nodeLevel]
//	node.next = levelNodeList.head.next
//	node.pre = levelNodeList.head
//	levelNodeList.head.next.pre = node
//	levelNodeList.head.next = node
//	this.cache[node.key] = node
//	this.size++
//	if this.size > this.cap {
//		this.removeTail()
//	}
//}
//func (this *LFUCache) removeTail() {
//	node := this.levelDoubleNodeList[this.minLevel].tail.pre
//	this.removeNode(node)
//}
//func (this *LFUCache) removeNode(node *DoubleNode) {
//	node.pre.next = node.next
//	node.next.pre = node.pre
//	delete(this.cache, node.key)
//	this.size--
//	//判断是不是minlevel
//	if node.level == this.minLevel && this.levelDoubleNodeList[node.level].head.next.next == nil {
//		this.minLevel++
//	}
//}
//func (this *LFUCache) initDoubleList() *DoubleNodeList {
//	head := &DoubleNode{key: 0, val: 0}
//	tail := &DoubleNode{key: 0, val: 0}
//	head.next = tail
//	tail.pre = head
//	return &DoubleNodeList{
//		head: head,
//		tail: tail,
//	}
//}
//
//func (this *LFUCache) Get(key int) int {
//	node := this.cache[key]
//	if node == nil {
//		return -1
//	}
//	this.removeNode(node)
//	node.level++
//	this.addNode(node)
//	return node.val
//}
//
//func (this *LFUCache) Put(key int, value int) {
//	node := this.cache[key]
//	if node == nil {
//		this.addNewNode(key, value)
//		return
//	}
//	node.val = value
//	this.Get(key)
//}

func Test_3(t *testing.T) {
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
