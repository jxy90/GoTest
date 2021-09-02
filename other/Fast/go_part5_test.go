package Fast

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func Test21(t *testing.T) {
	fmt.Println("----------")
	type location struct {
		Lat, Long float64
	}
	a := location{-4.1, 137.5}
	fmt.Printf("%v \n", a)
	fmt.Printf("%+v \n", a)
	fmt.Println("----------")
	b := a
	b.Lat = 1
	fmt.Printf("%+v \n", a)
	fmt.Printf("%+v \n", b)
	changeLocation := func(l location) {
		l.Lat += 100
		fmt.Printf("%+v \n", l)
	}
	changeLocation(a)
	fmt.Printf("%+v \n", a)
	fmt.Println("----------")
	//注意下小写
	bytes, _ := json.Marshal(b)
	fmt.Println(string(bytes))
}

type talker interface {
	talk() string
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew", int(l))
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type starship struct {
	laser
}

func Test23(t *testing.T) {
	//组合和嵌入
	type A struct {
		a int
	}
	type B struct {
		b int
	}
	type C struct {
		A A
		B B
	}
	type D struct {
		A
		B
	}
	fmt.Printf("%v \n", new(C).A.a)
	fmt.Printf("%v \n", new(D).a)
	fmt.Println("----------")
	s := starship{laser(3)}
	fmt.Println(s.talk())
	shout(s)
	fmt.Println("----------")
	fmt.Println("----------")
}
