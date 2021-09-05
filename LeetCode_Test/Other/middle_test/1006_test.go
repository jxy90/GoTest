package middle_test

import (
	"fmt"
	"testing"
)

func Test_clumsy(t *testing.T) {
	fmt.Println(clumsy(4))
	fmt.Println(clumsy(10))
}

func clumsy(N int) int {
	return clumsyHelper(N, 1)
}

// 递归函数，求n的笨阶乘
// prefix表示：前缀（符号位：加或减），取值：1 或 -1，分别表示：+ 或 -
// 最开始主函数调用时，符号位为加运算+，后面都是减运算-
func clumsyHelper(N int, prefix int) int {
	if N <= 2 {
		return N * prefix
	} else if N == 3 {
		return 6 * prefix
	} else {
		// 递归分解：先计算 n * (n-1) / (n-2) + (n-3)，再递归后面的
		return prefix*N*(N-1)/(N-2) + (N - 3) + clumsyHelper(N-4, -1) // prefix后面都是-1，表示先做减法
	}
}

func clumsy2(N int) int {
	options := []byte{'*', '/', '+', '-'}
	nums := make([]int, 0)
	ops := make([]byte, 0)
	// 维护运算符优先级
	priority := map[byte]int{}
	priority['*'] = 2
	priority['/'] = 2
	priority['+'] = 1
	priority['-'] = 1

	var calc func()
	calc = func() {
		b := nums[len(nums)-1]
		nums = nums[:len(nums)-1]

		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]

		a := nums[len(nums)-1]
		nums = nums[:len(nums)-1]

		ans := 0
		if op == '+' {
			ans = a + b
		} else if op == '-' {
			ans = a - b
		} else if op == '*' {
			ans = a * b
		} else if op == '/' {
			ans = a / b
		}
		nums = append(nums, ans)
	}

	index := 0
	for N > 0 {
		op := options[index%4]
		nums = append(nums, N)
		for len(ops) != 0 && priority[ops[len(ops)-1]] >= priority[op] {
			//计算
			calc()
		}
		if N != 1 {
			ops = append(ops, op)
		}
		N--
		index++
	}

	for len(ops) > 0 {
		//计算
		calc()
	}
	return nums[len(nums)-1]
}
