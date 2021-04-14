package middle_test

import "testing"

type Trie struct {
	isWord   bool
	children [26]*Trie
}

func ConstructorTrie() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	for i := range word {
		index := word[i] - 'a'
		if this.children[index] == nil {
			this.children[index] = &Trie{}
		}
		this = this.children[index]
	}
	this.isWord = true
}

func (this *Trie) Search(word string) bool {
	for i := range word {
		index := word[i] - 'a'
		if this.children[index] == nil {
			return false
		}
		this = this.children[index]
	}
	return this.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	for i := range prefix {
		index := prefix[i] - 'a'
		if this.children[index] == nil {
			return false
		}
		this = this.children[index]
	}
	return true
}

func Test_Tire(t *testing.T) {
	var obj Trie
	f := []string{"Trie", "insert", "insert", "insert", "insert", "insert", "insert", "search", "search", "search", "search", "search", "search", "search", "search", "search", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith"}
	m := []string{"", "app", "apple", "beer", "add", "jam", "rental", "apps", "app", "ad", "applepie", "rest", "jan", "rent", "beer", "jam", "apps", "app", "ad", "applepie", "rest", "jan", "rent", "beer", "jam"}
	ans := make([]bool, 0)
	for i := range f {
		method := f[i]
		word := m[i]
		if method == "Trie" {
			obj = ConstructorTrie()
		} else if method == "insert" {
			obj.Insert(word)
		} else if method == "search" {
			res := obj.Search(word)
			ans = append(ans, res)
		} else if method == "startsWith" {
			res := obj.StartsWith(word)
			ans = append(ans, res)
		}
	}
	print(ans)
}
