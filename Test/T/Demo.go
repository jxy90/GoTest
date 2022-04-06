package main

import (
	"fmt"
)

// Print 泛型
func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
	fmt.Println()
}
func Print2[T ~int | ~int32 | ~int64](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
	fmt.Println()
}

type People struct {
	name string
	age  int
}

func (p People) String() string {
	return fmt.Sprintf("{name:%v age: %v} ", p.name, p.age)
}

func main() {
	Print([]string{"你好, ", "脑子进了煎鱼\n"})
	Print([]int64{1, 2, 3})
	Print([]People{{"1", 1}, {"2", 2}})
}

func Map[T1, T2 interface{}](s []T1, f func(T1) T2) []T2 {
	res := make([]T2, len(s))
	for i, v := range s {
		res[i] = f(v)
	}
	return res
}
