package transfer

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Student struct {
	ID     int
	Gender string
	Name   string
	Sno    string
}
type Student2 struct {
	ID     int
	Gender string
	Name   string
	Sno    string
}

func Test_transfer(t *testing.T) {
	var s1 = Student{
		ID:     12,
		Gender: "男",
		Name:   "李四",
		Sno:    "s001",
	}
	fmt.Printf("%#v\n", s1)

	jsonByte, _ := json.Marshal(s1)
	//jsonStr := string(jsonByte)
	//fmt.Printf("%v", jsonStr)
	var s2 Student2
	err := json.Unmarshal(jsonByte, &s2)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", s2)
	fmt.Println(s2.Name)
}

func Test_transfer0(t *testing.T) {
	a := A{
		a: "A",
		b: 1,
		c: true,
	}
	transfer(a)
}

type A struct {
	a string
	b int
	c bool
}
type B struct {
	a string
	b int
	c bool
}

func transfer(i interface{}) {
	fmt.Printf("%v ", i.(B))
}
