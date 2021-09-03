package middle_test

import (
	"fmt"
	"math"
	"math/bits"
	"strconv"
	"testing"
)

func Test_smallestGoodBase(t *testing.T) {
	fmt.Println(smallestGoodBase("13"))
}

func smallestGoodBase(n string) string {
	v, _ := strconv.Atoi(n)
	//位数最长就是二进制情况下的位数长度
	lenmax := bits.Len(uint(v))
	for i := lenmax - 1; i > 1; i-- {
		k := int(math.Pow(float64(v), 1/float64(i)))
		sum, s := 1, 1
		for j := 0; j < i; j++ {
			s *= k
			sum += s
		}
		if sum == v {
			return strconv.Itoa(k)
		}
	}
	//这种情况下就是位数为2，一个1另一个就是v-1，直接返回v-1进制
	return strconv.Itoa(v - 1)
}

func smallestGoodBase0(n string) string {
	m, _ := strconv.ParseInt(n, 10, 64)
	max := 64
	for i := max; i >= 3; i-- {
		k := int64(math.Pow(float64(m), float64(1/(i-1))))
		res := int64(0)
		for j := 0; j < i; j++ {
			res = res*k + 1
		}
		if res == m {
			return fmt.Sprintf("%v", k)
		}
	}
	return fmt.Sprintf("%v", m-1)
}
