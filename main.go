package main

import (
	"fmt"
)

type People struct {
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
	p1 := People{"jxy1", 1}
	p2 := &People{"jxy2", 2}
	p1.getName()
	p1.getAge()
	p2.getName()
	p2.getAge()
}
