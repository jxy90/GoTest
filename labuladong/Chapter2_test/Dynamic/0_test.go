package Dynamic

import "testing"

func fib(n int) int {
	return fib(n-1) + fib(n-2)
}

func Test_0(t *testing.T) {

}
