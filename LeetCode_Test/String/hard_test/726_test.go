package hard_test

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
	"unicode"
)

func Test_countOfAtoms(t *testing.T) {
	//fmt.Println(countOfAtoms("H2O"))
	//fmt.Println(countOfAtoms("Mg(OH)2"))
	fmt.Println(countOfAtoms("H11He49NO35B7N46Li20"))

}

func countOfAtoms(formula string) string {
	n := len(formula)
	index := 0
	//得到数字
	getNum := func() int {
		if index == n || !unicode.IsDigit(rune(formula[index])) {
			return 1
		}
		num := 0
		for ; index < n && unicode.IsDigit(rune(formula[index])); index++ {
			num = num*10 + int(formula[index]-'0')
		}
		return num
	}
	//得到原子⚛
	getAtom := func() string {
		start := index
		index++
		for index < n && unicode.IsLower(rune(formula[index])) {
			index++
		}
		return formula[start:index]
	}
	//保存每层原子数
	stack := []map[string]int{{}}
	for index < n {
		if formula[index] == '(' {
			//入栈重新计算
			index++
			stack = append(stack, map[string]int{})
		} else if formula[index] == ')' {
			//出栈求和
			index++
			num := getNum()
			now := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for k, v := range now {
				stack[len(stack)-1][k] += v * num
			}
		} else {
			//记录当前数量
			atom := getAtom()
			num := getNum()
			stack[len(stack)-1][atom] += num
		}
	}
	//最后的map
	last := stack[0]
	//定义node方便排序
	type Node struct {
		Key string
		Val int
	}
	list := make([]*Node, 0, len(last))
	for k, v := range last {
		list = append(list, &Node{
			Key: k,
			Val: v,
		})
	}
	//排序
	sort.Slice(list, func(i, j int) bool {
		return list[i].Key < list[j].Key
	})
	//结果
	ans := ""
	for _, node := range list {
		ans += node.Key
		if node.Val > 1 {
			ans += strconv.Itoa(node.Val)
		}
	}
	return ans
}
