package middle_test

import (
	"fmt"
	"testing"
)

func Test_numRabbits(t *testing.T) {
	fmt.Println(numRabbits([]int{1, 0, 1, 0, 0}))
	fmt.Println(numRabbits([]int{1, 1, 2}))
}

//4ms 2.8MB
func numRabbits(answers []int) int {
	//保存爆出自己颜色的兔子有多少只
	hash := make(map[int]int, 0)
	for _, v := range answers {
		//加一的作用是加上自己的数量
		hash[v+1]++
	}
	ans := 0
	for key := range hash {
		if key == 0 {
			ans += hash[key]
		} else {
			group := (hash[key]-1)/key + 1
			ans += key * group
		}
	}
	return ans
}
