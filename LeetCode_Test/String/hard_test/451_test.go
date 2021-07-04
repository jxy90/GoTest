package hard_test

import (
	"sort"
	"testing"
)

func Test_frequencySort(t *testing.T) {
	println(frequencySort("tree"))
}

type ByteNode struct {
	Key   byte
	Count int
}

func frequencySort(s string) string {
	cache := map[byte]*ByteNode{}
	array := make([]*ByteNode, 0)
	for i := range s {
		if _, ok := cache[s[i]]; ok {
			node := cache[s[i]]
			node.Count++
		} else {
			node := &ByteNode{
				Key:   s[i],
				Count: 1,
			}
			cache[s[i]] = node
			array = append(array, node)
		}
	}
	sort.Slice(array, func(i, j int) bool {
		return array[i].Count > array[j].Count
	})
	ans := ""
	for _, node := range array {
		for node.Count > 0 {
			ans += string(node.Key)
			node.Count--
		}
	}
	return ans
}
