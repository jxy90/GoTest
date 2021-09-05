package middle_test

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_rand10(t *testing.T) {
	fmt.Println(rand10())
}

func rand10() int {
	for {
		row := rand7()
		col := rand7()
		idx := (row-1)*7 + col
		if idx <= 40 {
			return 1 + (idx-1)%10
		}
	}
}

func rand7() int {
	return rand.Intn(7)
}
