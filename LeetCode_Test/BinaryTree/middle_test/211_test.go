package middle_test

import (
	"fmt"
	"testing"
)

func Test_WordDictionary(t *testing.T) {
	wordDictionary := ConstructorWordDictionary()
	fmt.Println(wordDictionary.Search("a")) // return False
	//wordDictionary.AddWord("bad")
	//wordDictionary.AddWord("dad")
	//wordDictionary.AddWord("mad")
	//fmt.Println(wordDictionary.Search("pad")) // return False
	//fmt.Println(wordDictionary.Search("bad")) // return True
	//fmt.Println(wordDictionary.Search(".ad")) // return True
	//fmt.Println(wordDictionary.Search("b..")) // return True

}

type WordDictionary struct {
	isWord   bool
	children [26]*WordDictionary
}

func ConstructorWordDictionary() WordDictionary {
	return WordDictionary{
		isWord:   false,
		children: [26]*WordDictionary{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	for _, char := range word {
		index := char - 'a'
		if this.children[index] == nil {
			this.children[index] = &WordDictionary{
				isWord:   false,
				children: [26]*WordDictionary{},
			}
		}
		this = this.children[index]
	}
	this.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.SearchIndex(word, 0)
}
func (this *WordDictionary) SearchIndex(word string, cur int) bool {
	index := int(word[cur] - 'a')
	if cur == len(word)-1 {
		if index == 205 {
			for _, child := range this.children {
				if child != nil && child.isWord {
					return true
				}
			}
			return false
		}
		if this.children[index] == nil {
			return false
		}
		return this.children[index].isWord
	}
	if index == 205 {
		for _, child := range this.children {
			if child != nil && child.SearchIndex(word, cur+1) {
				return true
			}
		}
		return false
	}
	if this.children[index] == nil {
		return false
	}
	return this.children[index].SearchIndex(word, cur+1)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
