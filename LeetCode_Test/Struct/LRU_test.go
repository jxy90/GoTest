package Struct_test

import (
	"testing"
)

func Test_LRUCache(t *testing.T) {
	methods := []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
	params := [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}

	obj := ConstructorLRU(params[0][0])
	for i := 1; i < len(methods); i++ {
		if methods[i] == "put" {
			obj.Put(params[i][0], params[i][1])
		} else if methods[i] == "get" {
			obj.Get(params[i][0])
		}
	}
	fmt.Println(obj.tail.value)
}

type DLinkNode struct {
	key, value int
	pre, next  *DLinkNode
}

type LRUCache struct {
	cache      map[int]*DLinkNode
	cap        int
	head, tail *DLinkNode
}

func ConstructorLRU(capacity int) LRUCache {
	lru := LRUCache{
		cache: map[int]*DLinkNode{},
		cap:   capacity,
		head:  &DLinkNode{},
		tail:  &DLinkNode{},
	}
	lru.head.next = lru.tail
	lru.tail.pre = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	node := this.cache[key]
	if node == nil {
		return -1
	}
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node := this.cache[key]
	if node != nil {
		node.value = value
		this.moveToHead(node)
		return
	}
	newNode := &DLinkNode{key: key, value: value}
	this.addToHead(newNode)
	if len(this.cache) > this.cap {
		this.removeTail()
	}
}

func (this *LRUCache) moveToHead(node *DLinkNode) {
	this.removeNode(node)
	this.addToHead(node)
}
func (this *LRUCache) removeNode(node *DLinkNode) {
	node.next.pre = node.pre
	node.pre.next = node.next
	delete(this.cache, node.key)
}
func (this *LRUCache) addToHead(node *DLinkNode) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	this.head.next = node
	this.cache[node.key] = node
}
func (this *LRUCache) removeTail() *DLinkNode {
	node := this.tail.pre
	this.removeNode(node)
	return node
}
