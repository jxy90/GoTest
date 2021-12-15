package middle_test

import (
	"fmt"
	"testing"
)

func Test_loudAndRich(t *testing.T) {
	fmt.Println(loudAndRich([][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}}, []int{3, 2, 5, 4, 6, 1, 7, 0}))
}

func loudAndRich(richer [][]int, quiet []int) []int {
	//统计比我有钱的人到cache
	cache := make(map[int][]int)
	for _, ints := range richer {
		a, b := ints[0], ints[1]
		cache[b] = append(cache[b], a)
	}
	//初始化返回值,都是-1表示未被赋值
	n := len(quiet)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	var dfs func(person int)
	dfs = func(person int) {
		//如果!=-1,说明已找到最安静的人
		if ans[person] != -1 {
			return
		}
		//先把自己设成最安静的人
		ans[person] = person
		//遍历比我有钱的人
		for _, next := range cache[person] {
			//去找比我有钱的人的最安静的人
			dfs(next)
			//对比找对我自己最安静的人
			if quiet[ans[next]] < quiet[ans[person]] {
				ans[person] = ans[next]
			}
		}
	}
	//对每一个人进行寻找
	for i := range ans {
		dfs(i)
	}
	return ans
}
