package old

import (
	"fmt"
	"testing"
)

func Test_0(t *testing.T) {
	fmt.Println("")
	dog := Dog{
		Animal:   Animal{"dog", 1},
		BoomType: "D",
	}
	dog.Print()
}

type Animal struct {
	Name string
	Age  int
}

func (a *Animal) Print() {
	fmt.Printf("Name:%v,Age:%v", a.Name, a.Age)
}

type Dog struct {
	Animal
	BoomType string
}
