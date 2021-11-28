package middle_test

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_Solution519(t *testing.T) {
	fmt.Println(Constructor519(3, 1))
}

type Solution519 struct {
	nums        map[int]int
	m, n, total int
}

func Constructor519(m int, n int) Solution519 {
	return Solution519{
		nums:  map[int]int{},
		m:     m,
		n:     n,
		total: m * n,
	}
}

func (this *Solution519) Flip() (ans []int) {
	randNum := rand.Intn(this.total)
	this.total--
	//当前randnum有没有被访问过,
	//有的话,取当时交换过末位的值,
	//没有的话,取当前值.
	if x, ok := this.nums[randNum]; ok {
		ans = []int{x / this.n, x % this.n}
	} else {
		ans = []int{randNum / this.n, randNum % this.n}
	}
	//在randNum处记录,未被访问的数
	if x, ok := this.nums[this.total]; ok {
		this.nums[randNum] = x
	} else {
		this.nums[randNum] = this.total
	}
	return
}

func (this *Solution519) Reset() {
	this.total = this.m * this.n
	this.nums = map[int]int{}
}
