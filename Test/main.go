package main

import (
	"fmt"
	"sort"
)

func sum(a int, b int) int {
	return a + b
}

type A struct {
	nums []int
	num  int
	str  string
}

func main() {
	//data := []int{10, 25, 11, 24, 17, 26}
	//i := sort.Search(len(data), func(i int) bool {
	//	return data[i] >= 23
	//})
	//fmt.Println("最终的结果为", i)
	//data := []int{10, 22, 11, 22, 17, 26}
	data := []int{1, 10, 4, 4, 2, 7}
	rec := append(sort.IntSlice(nil), data...)
	rec.Sort()
	rec2 := sort.IntSlice(data)
	rec2.Sort()
	fmt.Println("最终的结果为", rec)
	fmt.Println("最终的结果为", rec2)
	fmt.Println("hello test")
	fmt.Println('.' - 'a')

	a := new(A)
	b := A{}
	fmt.Printf("a:%v b:%v", a, b)
}
