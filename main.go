package main

import (
	"fmt"
)

type People struct {
	name  string
	age   int
	hobby []string
}
type Animal struct {
	name string
	age  int
}

func (p People) getName() {
	fmt.Println(p.name)
}

func (p *People) getAge() {
	fmt.Println(p.age)
}

func main() {
	p1 := People{"jxy1", 1, nil}
	p2 := &People{"jxy2", 2, nil}
	p1.getName()
	p1.getAge()
	p2.getName()
	p2.getAge()
	s1 := make([]int, 10)
	m1 := map[string]int{}
	fmt.Println(s1)
	fmt.Println(m1)
	fmt.Println(p1)
	change(s1, m1, &p1)
	fmt.Println(s1)
	fmt.Println(m1)
	fmt.Println(p1)
}

func change(s1 []int, m1 map[string]int, people *People) {
	s1[0] = 1
	m1["1"] = 1
	people.age = 2
	people.hobby = []string{"1"}
}
