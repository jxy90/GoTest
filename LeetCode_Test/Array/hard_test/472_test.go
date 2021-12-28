package hard_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_findAllConcatenatedWordsInADict(t *testing.T) {
	fmt.Println(findAllConcatenatedWordsInADict([]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}))
}

type Tire struct {
	isWild   bool
	children [26]*Tire
}

func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	root := &Tire{}
	ans := make([]string, 0)
	for _, word := range words {
		if word == "" {
			continue
		}
		if !root.dfs(word) {
			root.insert(word)
		} else {
			ans = append(ans, word)
		}
	}
	return ans
}

func (root *Tire) dfs(word string) bool {
	if word == "" {
		return true
	}
	node := root
	for i, ch := range word {
		node = node.children[ch-'a']
		if node == nil {
			return false
		}
		if node.isWild && root.dfs(word[i+1:]) {
			return true
		}
	}
	return false
}

func (t *Tire) insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Tire{}
		}
		node = node.children[ch]
	}
	node.isWild = true
}
