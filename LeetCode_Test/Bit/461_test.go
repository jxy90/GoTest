package Bit

import "testing"

func Test_hammingDistance(t *testing.T) {
	println(hammingDistance(1, 4))
	println(hammingDistance(93, 73))
}

func hammingDistance(x int, y int) int {
	count := 0
	temp := x ^ y
	for temp != 0 {
		if temp&1 == 1 {
			count++
		}
		temp >>= 1
	}
	return count
}

func hammingDistance1(x int, y int) int {
	count := 0
	for x|y != 0 {
		a, b := x&1, y&1
		if a != b {
			count++
		}
		x >>= 1
		y >>= 1
	}
	return count
}

func hammingDistance0(x int, y int) int {
	n := 31
	count := 0
	for i := 0; i < n; i++ {
		if x>>i&1 != y>>i&1 {
			count++
		}
	}
	return count
}
