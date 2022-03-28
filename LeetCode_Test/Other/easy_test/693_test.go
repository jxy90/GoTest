package easy_test

import (
	"fmt"
	"log"
	"testing"
)

func Test_693(t *testing.T) {
	fmt.Printf("%b \r\n", 5&2)
	log.Println(hasAlternatingBits(5))
	log.Println(hasAlternatingBits(7))
	log.Println(hasAlternatingBits(11))
}

func hasAlternatingBits(n int) bool {
	m := n >> 1
	temp := n ^ m
	//fmt.Printf("%v(%b)&%v(%b)=%b \r\n", n, n, m, m, temp)
	return temp&(temp+1) == 0
}
