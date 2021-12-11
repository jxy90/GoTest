package Struct

import "sort"

type Node struct {
	value     string
	timestamp int
}

type TimeMap struct {
	m map[string][]Node
}

/** Initialize your data structure here. */
func ConstructorTimeMap() TimeMap {
	return TimeMap{map[string][]Node{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], Node{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	Nodes := this.m[key]
	index := sort.Search(len(Nodes), func(i int) bool {
		return Nodes[i].timestamp > timestamp
	})
	if index > 0 {
		return Nodes[index-1].value
	}
	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := ConstructorMapSum();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
