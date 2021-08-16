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
	println(1)
}

type NodeList []*Node
type Node struct {
	prev, next      *Node
	key, val, level int
}

type LFUCache struct {
	minLevel, capacity int
	cache              map[int]*Node
	levelCache         map[int]NodeList
}

func ConstructorLFU(capacity int) LFUCache {
	lfu := LFUCache{
		minLevel:   0,
		capacity:   capacity,
		cache:      map[int]*Node{},
		levelCache: map[int]NodeList{},
	}
	return lfu
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.removeNode(node)
		if len(this.levelCache[node.level]) == 0 && node.level == this.minLevel {
			this.minLevel++
		}
		node.level++
		this.addNode(node)
		return node.val
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.Get(key)
	} else {
		node := &Node{
			key:   key,
			val:   value,
			level: 0,
		}
		this.addNode(node)
		if len(this.cache) > this.capacity {
			//容量超了删除尾部
			this.removeTail()
		}
		//一定要最后更新
		this.minLevel = 0
	}
}

func (this *LFUCache) removeTail() {
	list := this.levelCache[this.minLevel]
	n := len(list)
	removeNode := list[n-1]
	//缓存中移除
	delete(this.cache, removeNode.key)
	//移除尾部
	list = list[:len(list)-1]
	this.levelCache[this.minLevel] = list
}
func (this *LFUCache) addNode(node *Node) {
	this.cache[node.key] = node
	nodeList := this.levelCache[node.level].addNode(node)
	this.levelCache[node.level] = nodeList
}

func (nodeList NodeList) addNode(node *Node) NodeList {
	newList := make(NodeList, 0)
	newList = append(newList, node)
	newList = append(newList, nodeList...)
	return newList
}
func (this *LFUCache) removeNode(node *Node) {
	delete(this.cache, node.key)
	list := this.levelCache[node.level]
	for i, n := range list {
		if node == n {
			list = append((list)[:i], (list)[i+1:]...)
		}
	}
	this.levelCache[node.level] = list
}
