package main

import (
	"fmt"
	"sync"
	"testing"
)

func mySqrt(x int) int {
	ans := 0
	l, r := 0, x
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			l = mid + 1
			ans = mid
		} else {
			r = mid - 1
		}
	}
	return ans
}

func Test_mySqrt(t *testing.T) {
	var a = 3
	var b = 4
	res := sum(a, b)
	fmt.Printf("%d 与%d之和:为%d", a, b, res)
	if res != 7 {
		t.Error("error")
	}

	fmt.Println(mySqrt(4))
}

/*func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	ans := myPowHelper(x, n)
	return ans
}

func myPowHelper(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	return x * myPowHelper(x, n-1)
}*/
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	ans := myPowHelper(x, n)
	return ans
}

func myPowHelper(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	if n%2 == 0 {
		half := myPowHelper(x, n/2)
		return half * half
	} else {
		half := myPowHelper(x, n/2)
		return half * half * x
	}
}
func Test_myPow(t *testing.T) {
	//fmt.Println(myPow(0.00001, 2147483647))
	//fmt.Println(math.Pow(0.00001, 2147483647))
	channel()
}

func channel() {
	c := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			fmt.Println(<-c)
			wg.Done()
		}
	}()
	c <- true
	wg.Wait()
}
