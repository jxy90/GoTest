package hard_test

import (
	"fmt"
	"math"
	"testing"
)

func Test_poorPigs(t *testing.T) {
	//fmt.Println(maxSumSubmatrix([][]int{{1, 0, 1}, {0, -2, 3}}, 2))
	fmt.Println(poorPigs(1000, 15, 60))
	fmt.Println(poorPigs(4, 15, 15))
	fmt.Println(poorPigs(4, 15, 30))
}

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	//cnt一只猪,可以检测的信息量
	cnt := minutesToTest/minutesToDie + 1
	//bucketsOnce := buckets / cnt
	//fmt.Println(bucketsOnce)
	//pow(base,ans)>=buckets
	//base的ans次方,大于等于 桶数量
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(cnt))))
}
