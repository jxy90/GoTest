package easy_test

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_reverseBits(t *testing.T) {
	//4
	fmt.Println(reverseBits2(-1))
	fmt.Println(reverseBits2(7))
	//8
	fmt.Println(reverseBits2(1775))
	//31
	fmt.Println(reverseBits2(2147483647))
	//32
	fmt.Println(reverseBits2(2147483646))
}

func reverseBits2(num int) int {
	counts := make([]int, 0)
	count := 0
	index := 0
	for num != 0 && index < 32 {
		if num&1 == 1 {
			count++
		} else {
			counts = append(counts, count)
			count = 0
			counts = append(counts, 0)
		}
		num >>= 1
		index++
	}
	counts = append(counts, count)
	counts = append(counts, 0)
	max := 0
	if len(counts) < 3 {
		max = counts[0] + counts[1] + 1
	}
	for i := 2; i < len(counts); i++ {
		max = CommonUtil.Max(max, counts[i]+counts[i-2]+1)
	}
	if max > 32 {
		return 32
	}
	return max
}
