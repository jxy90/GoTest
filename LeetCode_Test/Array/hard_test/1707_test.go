package hard_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"math"
	"sort"
	"testing"
)

func Test_maximizeXor(t *testing.T) {
	//fmt.Println(minimumTimeRequired([]int{3, 2, 3}, 3))
	fmt.Println(maximizeXor([]int{0, 1, 2, 3, 4}, [][]int{{3, 1}, {1, 3}, {5, 6}}))
	fmt.Println(maximizeXor([]int{5, 2, 4, 6, 6, 3}, [][]int{{12, 4}, {8, 1}, {6, 3}}))
}

func maximizeXor(nums []int, queries [][]int) []int {
	trie := &Trie{
		min: math.MaxInt32,
	}
	for i := range nums {
		trie.insert(nums[i])
	}
	ans := make([]int, len(queries))
	for i, query := range queries {
		ans[i] = trie.getMaxXor(query[0], query[1])
	}
	return ans
}

const L = 30

type Trie struct {
	children [2]*Trie
	min      int
}

func (t *Trie) insert(val int) {
	node := t
	if val < node.min {
		node.min = val
	}
	for i := L - 1; i >= 0; i-- {
		bit := val >> i & 1
		if node.children[bit] == nil {
			node.children[bit] = &Trie{
				min: val,
			}
		}
		node = node.children[bit]
		if val < node.min {
			node.min = val
		}
	}
}

func (t *Trie) getMaxXor(val, limit int) int {
	node := t
	if node.min > limit {
		return -1
	}
	ans := 0
	for i := L - 1; i >= 0; i-- {
		bit := val >> i & 1
		if node.children[bit^1] != nil && node.children[bit^1].min <= limit {
			bit ^= 1
			ans |= 1 << i
		}
		node = node.children[bit]
	}

	return ans
}

func maximizeXor0(nums []int, queries [][]int) []int {
	sort.Ints(nums)
	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	ans := make([]int, len(queries))
	for _, query := range queries {
		temp := -1
		for _, num := range nums {
			if num > query[1] {
				break
			}
			temp = CommonUtil.Max(temp, num^query[0])
		}
	}
	return ans
}
