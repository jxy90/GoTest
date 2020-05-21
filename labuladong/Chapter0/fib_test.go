package Chapter0_test

import "testing"

//1.暴力递归
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

//2。带备忘录的递归解法
func fibHis(n int) int {
	his := map[int]int{}
	his[1] = 1
	his[2] = 1
	return fibHisHelper(his, n)
}

func fibHisHelper(his map[int]int, n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	if his[n] != 0 {
		return his[n]
	}
	his[n] = fibHisHelper(his, n-1) + fibHisHelper(his, n-2)
	return his[n]
}

//3。dp数组的迭代解法
func fibdp(n int) int {
	his := map[int]int{}
	his[1] = 1
	his[2] = 1
	for i := 3; i <= n; i++ {
		his[i] = his[i-1] + his[i-2]
	}
	return his[n]
}

//4.dp优化不使用数组进行保存，降低空间复杂度
func fibdp2(n int ) int {
	pre,cur :=1,1
	for i:=3;i<=n;i++ {
		sum :=pre+cur
		pre=cur
		cur=sum
	}
	return cur
}

func Test_fib(t *testing.T) {
	println(fib(10))
	println(fibHis(10))
	println(fibdp(10))
	println(fibdp2(10))
}
