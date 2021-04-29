package middle_test

import "testing"

func Test_LRUCache(t *testing.T) {
	methods := []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
	params := [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}

	obj := Constructor(params[0][0])
	for i := 1; i < len(methods); i++ {
		if methods[i] == "put" {
			obj.Put(params[i][0], params[i][1])
		} else if methods[i] == "get" {
			obj.Get(params[i][0])
		}
	}
	println(obj.tail.value)
}

type DLinkedNode struct {
	key, value int
	pre, next  *DLinkedNode
}

type LRUCache struct {
	cache      map[int]*DLinkedNode
	capacity   int
	head, tail *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cache:    map[int]*DLinkedNode{},
		capacity: capacity,
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}
	lru.head.next = lru.tail
	lru.tail.pre = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	node := this.cache[key]
	//不存在返回-1
	if node == nil {
		return -1
	}
	//存在 移动到头部
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node := this.cache[key]
	//存在 更新val,移动到头部
	if node != nil {
		node.value = value
		this.moveToHead(node)
		return
	}
	//不存在
	newNode := initDLinkedNode(key, value)
	this.addToHead(newNode)
	this.cache[key] = newNode
	//缓存满了 删除
	if len(this.cache) > this.capacity {
		remove := this.removeTail()
		delete(this.cache, remove.key)
	}
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.pre
	this.removeNode(node)
	return node
}
func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}
func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}

///**
// * Your LRUCache object will be instantiated and called as such:
// * obj := Constructor(capacity);
// * param_1 := obj.Get(key);
// * obj.Put(key,value);
// */
