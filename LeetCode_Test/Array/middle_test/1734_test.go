package middle_test

import (
	"testing"
)

func Test_decode(t *testing.T) {
	fmt.Println(decode([]int{3, 1}))
	fmt.Println(decode([]int{6, 5, 4, 6}))
}

func decode(encoded []int) []int {
	n := len(encoded)
	//1~n的异或
	totalXOR0 := 0
	for i := 1; i <= n+1; i++ {
		totalXOR0 ^= i
	}
	//encoded[0] = perm[0] ^ perm[1]
	//encoded[1] = perm[1] ^ perm[2]
	//encoded[2] = perm[2] ^ perm[3]
	//encoded[3] = perm[3] ^ perm[4]
	//encoded[4] = perm[4] ^ perm[5]
	//encoded[1]~encoded[n-1]的异或
	totalXOR1 := 0
	for i := 1; i < n; i += 2 {
		totalXOR1 ^= encoded[i]
	}
	perm0 := totalXOR0 ^ totalXOR1
	perm := make([]int, 0)
	perm = append(perm, perm0)
	temp := perm0
	for i := range encoded {
		temp = encoded[i] ^ temp
		perm = append(perm, temp)
	}
	return perm
}
