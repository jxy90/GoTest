package unionFindDisjointSets

import (
	"fmt"
	"sort"
	"testing"
)

func Test_accountsMerge(t *testing.T) {
	fmt.Println(accountsMerge([][]string{{"John", "johnsmith@mail.com", "john00@mail.com"}, {"John", "johnnybravo@mail.com"}, {"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"Mary", "mary@mail.com"}}))
}

func accountsMerge(accounts [][]string) (ans [][]string) {
	emailToIndex := map[string]int{}
	emailToName := map[string]string{}
	for _, account := range accounts {
		name := account[0]
		for _, email := range account[1:] {
			if _, has := emailToIndex[email]; !has {
				emailToIndex[email] = len(emailToIndex)
				emailToName[email] = name
			}
		}
	}

	parent := make([]int, len(emailToIndex))
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(from, to int) {
		parent[find(from)] = find(to)
	}

	for _, account := range accounts {
		firstIndex := emailToIndex[account[1]]
		for _, email := range account[2:] {
			union(emailToIndex[email], firstIndex)
		}
	}

	indexToEmails := map[int][]string{}
	for email, index := range emailToIndex {
		index = find(index)
		indexToEmails[index] = append(indexToEmails[index], email)
	}

	for _, emails := range indexToEmails {
		sort.Strings(emails)
		account := append([]string{emailToName[emails[0]]}, emails...)
		ans = append(ans, account)
	}
	return
}

func accountsMerge0(accounts [][]string) [][]string {
	n := len(accounts)
	union := ConstructorUnionFind721(n)
	//cache := make(map[int]string)
	cache := make([]map[string]bool, n)
	for i := 0; i < n; i++ {
		for j := range accounts[i] {
			if cache[i] == nil {
				cache[i] = make(map[string]bool)
			}
			cache[i][accounts[i][j]] = true
		}
	}
	for i := 0; i < n; i++ {
		for j := range accounts[i] {
			if j == 0 {
				continue
			}
			for k := range cache {
				if k == i {
					continue
				}
				if cache[k][accounts[i][j]] {
					union.union(k, i)
					appendCache(accounts, cache, k, i)
				}
			}
		}
	}
	ans := make([][]string, 0)
	for i := 0; i < n; i++ {
		if union.parent[i] == i {
			ans = append(ans, accounts[i])
		}
	}
	return ans
}

func appendCache(accounts [][]string, cache []map[string]bool, k, i int) {
	for index := range accounts[k] {
		if cache[i][accounts[k][index]] == false {
			accounts[i] = append(accounts[i], accounts[k][index])
		}
	}
}

type UnionFind721 struct {
	parent []int
}

func ConstructorUnionFind721(total int) *UnionFind721 {
	p := make([]int, total)
	for i := 0; i < total; i++ {
		p[i] = i
	}
	return &UnionFind721{
		parent: p,
	}
}
func (u *UnionFind721) union(x, y int) {
	xp := u.find(x)
	yp := u.find(y)
	if xp == yp {
		return
	}
	u.parent[xp] = yp
}
func (u *UnionFind721) find(x int) int {
	root := x
	for root != u.parent[root] {
		root = u.parent[root]
	}
	i, j := x, 0
	for root != u.parent[i] {
		j = u.parent[i]
		u.parent[i] = root
		i = j
	}
	return root
}
