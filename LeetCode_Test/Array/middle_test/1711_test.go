package middle_test

import (
	"testing"
)

func Test_countPairs(t *testing.T) {
	println(countPairs([]int{1, 3, 5, 7, 9}))
	println(countPairs([]int{1, 1, 1, 3, 3, 3, 7}))
	println(countPairs([]int{2160, 1936, 3, 29, 27, 5, 2503, 1593, 2, 0, 16, 0, 3860, 28908, 6, 2, 15, 49, 6246, 1946, 23, 105, 7996, 196, 0, 2, 55, 457, 5, 3, 924, 7268, 16, 48, 4, 0, 12, 116, 2628, 1468}))
}

func countPairs(deliciousness []int) int {
	const mod = 1000000007
	//取数组最大值
	maxVal := deliciousness[0]
	for _, v := range deliciousness {
		if v > maxVal {
			maxVal = v
		}
	}
	//则两数之和不可能超过maxSum
	maxSum := maxVal * 2
	count := map[int]int{}
	ans := 0
	for _, v := range deliciousness {
		//左移符合条件的sum值,在count中得到符合条件的数量
		for sum := 1; sum <= maxSum; sum <<= 1 {
			ans += count[sum-v]
		}
		count[v]++
	}
	return ans % mod
}

func countPairs0(deliciousness []int) int {
	const mod = 1000000007
	//判断2的幂
	is2Power := func(num int) bool {
		return num > 0 && num&(num-1) == 0
	}

	n := len(deliciousness)
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			sum := deliciousness[i] + deliciousness[j]
			if is2Power(sum) {
				ans++
			}
		}
	}
	return ans % mod
}
