package Tree

import (
	"testing"
)

//type Trie struct {
//	children [26]*Trie
//	isWord   bool
//}
//
//func Constructor() Trie {
//	return Trie{}
//}
//
//func (this *Trie) Insert(word string) {
//	root := this
//	for i := range word {
//		if root.children[word[i]-'a'] == nil {
//			root.children[word[i]-'a'] = &Trie{}
//		}
//		root = root.children[word[i]-'a']
//	}
//	root.isWord = true
//}
//
//func (this *Trie) Search(word string) bool {
//	root := this
//	for i := range word {
//		if root.children[word[i]-'a'] == nil {
//			return false
//		}
//		root = root.children[word[i]-'a']
//	}
//	return root.isWord
//}
//
//func (this *Trie) StartsWith(prefix string) bool {
//	root := this
//	for i := range prefix {
//		if root.children[prefix[i]-'a'] == nil {
//			return false
//		}
//		root = root.children[prefix[i]-'a']
//	}
//	return true
//}

func Test_4(t *testing.T) {
}
