package Fast

import (
	"fmt"
	"testing"
)

func Test26(t *testing.T) {
	num := 42
	fmt.Println(&num)
	fmt.Println(*&num)
	fmt.Println("----------")
	address := &num
	fmt.Printf("address: %T\r\n", address)
	fmt.Println("----------")
	canada := "Canada"
	//声明指针类型*
	var home *string
	home = &canada
	//解引用
	fmt.Println(*home)
}

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func Test28_0(t *testing.T) {
	fmt.Println(increaseA())
	fmt.Println(increaseB())
	fmt.Println("----------")
}

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func f4() (r int) {
	t := 5
	defer func() {
		r = t + 5
	}()
	return t
}

func Test28_1(t *testing.T) {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println("----------")
}
