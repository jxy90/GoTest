package middle_test

import (
	"fmt"
	"testing"
)

func Test_2028(t *testing.T) {
	//fmt.Println(missingRolls([]int{3, 2, 4, 3}, 4, 2))
	fmt.Println(missingRolls([]int{1, 5, 6}, 3, 4))
}

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	//记录总点数,剩余点数
	total := (m + n) * mean
	left := total
	for _, roll := range rolls {
		left -= roll
	}
	//范围判断
	temp := float64(left) / float64(n)
	//fmt.Println(temp)
	if temp < float64(1) || temp > float64(6) {
		return []int{}
	}
	//平均分配
	tempInt := int(temp)
	yu := left % n
	ans := make([]int, n)
	for i := range ans {
		ans[i] = tempInt
		if yu > 0 {
			ans[i]++
		}
		yu--
	}
	return ans
}
