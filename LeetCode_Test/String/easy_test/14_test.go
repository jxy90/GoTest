package easy_test

import (
	"fmt"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	data := longestCommonPrefix(strs)
	fmt.Println(data)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	//把第一个string作为模式串
	for i := 0; i < len(strs[0]); i++ {
		//从1开始，后面的string，作为比较串
		for j := 1; j < len(strs); j++ {
			// 当模式串的长度>比较串的时候，判断相等。返回的是 比较串的本身
			// 当模式串和比较串相同位置不相等时，返回的是 开始位置到当前位置
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
