package middle_test

import (
	"sort"
)

func topKFrequent(words []string, k int) []string {
	cache := make(map[string]int)
	for i := range words {
		cache[words[i]]++
	}
	strs := make([]string, 0)
	for s := range cache {
		strs = append(strs, s)
	}
	sort.Slice(strs, func(i, j int) bool {
		a := strs[i]
		b := strs[j]
		return cache[a] > cache[b] || (cache[a] == cache[b] && a < b)
	})
	return strs[:k]
}
