package Struct_test

import "testing"

func Test_LFUCache(t *testing.T) {
	methods := []string{"LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"}
	params := [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {3}, {4, 4}, {1}, {3}, {4}}

	obj := ConstructorLFU(params[0][0])
	for i := 1; i < len(methods); i++ {
		if methods[i] == "put" {
			obj.Put(params[i][0], params[i][1])
		} else if methods[i] == "get" {
			obj.Get(params[i][0])
		}
	}
	fmt.Println(1)
}

type DoubleLinkNodeList struct {
	head, tail *DoubleLinkNode
}
type DoubleLinkNode struct {
	key, val, level int
	pre, next       *DoubleLinkNode
}

type LFUCache struct {
	cap, minLevel, size int
	cache               map[int]*DoubleLinkNode
	cacheLevelNodes     map[int]*DoubleLinkNodeList
}

func ConstructorLFU(capacity int) LFUCache {
	lfu := LFUCache{
		cap:             capacity,
		minLevel:        0,
		size:            0,
		cache:           map[int]*DoubleLinkNode{},
		cacheLevelNodes: map[int]*DoubleLinkNodeList{},
	}
	return lfu
}

func (this *LFUCache) Get(key int) int {
	//检查存在
	node := this.cache[key]
	if node == nil {
		return -1
	}
	//存在节点升级
	this.IncNodeLevel(node)
	return node.val
}

func (this *LFUCache) Put(key int, value int) {
	//检查存在
	node := this.cache[key]
	if node == nil {
		this.AddNewNode(key, value)
	} else {
		node.val = value
		this.Get(key)
	}
}

func (this *LFUCache) IncNodeLevel(node *DoubleLinkNode) {
	//移除老的
	this.removeNode(node)
	//升级
	node.level++
	//添加新的
	this.addNode(node)
}

func (this *LFUCache) AddNewNode(key int, value int) {
	node := &DoubleLinkNode{
		key:   key,
		val:   value,
		level: 0,
		pre:   nil,
		next:  nil,
	}
	//添加新的
	this.addNode(node)
	//重置minLevel
	this.minLevel = 0
}

func (this *LFUCache) removeNode(node *DoubleLinkNode) {
	node.next.pre = node.pre
	node.pre.next = node.next
	delete(this.cache, node.key)
	//节点总数减一
	this.size--
	//判断减去的节点是不是最小一行的最后一个
	if this.cacheLevelNodes[node.level].head.next.next == nil && node.level == this.minLevel {
		this.minLevel++
	}
}
func (this *LFUCache) addNode(node *DoubleLinkNode) {
	//放入cacheLevelNodes
	if this.cacheLevelNodes[node.level] == nil {
		this.initLevelNode(node)
	}
	levelNodes := this.cacheLevelNodes[node.level]
	node.next = levelNodes.head.next
	node.pre = levelNodes.head
	levelNodes.head.next.pre = node
	levelNodes.head.next = node
	this.size++
	//放入cache
	this.cache[node.key] = node
	if this.size > this.cap {
		this.removeTail()
	}
}

func (this *LFUCache) removeTail() {
	node := this.cacheLevelNodes[this.minLevel].tail.pre
	this.removeNode(node)
}

func (this *LFUCache) initLevelNode(node *DoubleLinkNode) {
	this.cacheLevelNodes[node.level] = &DoubleLinkNodeList{
		head: &DoubleLinkNode{},
		tail: &DoubleLinkNode{},
	}
	this.cacheLevelNodes[node.level].head.next = this.cacheLevelNodes[node.level].tail
	this.cacheLevelNodes[node.level].tail.pre = this.cacheLevelNodes[node.level].head
}
