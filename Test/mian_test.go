package main

import (
	"fmt"
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
	//if x==0 {
	//	return 0
	//}
	//i:= 1
	//for ;i<=x;i++{
	//	if i*i==x{
	//		return i
	//	}else if i*i>x {
	//		return i-1
	//	}
	//}
	//return i
}

func TestSum(t *testing.T) {
	var a = 3
	var b = 4
	res := sum(a, b)
	fmt.Printf("%d 与%d之和:为%d", a, b, res)
	if res != 7 {
		t.Error("error")
	}

	println(mySqrt(4))
}
