package Fast

import (
	"encoding/json"
	"fmt"
	"math/rand"
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

type pkq struct {
}
type mm struct {
}

type animaler interface {
	talk() string
	eat() string
}

func (p pkq) talk() string {
	return "pk pk "
}
func (p mm) talk() string {
	return "miao miao "
}
func (p pkq) eat() string {
	return "pk eat"
}
func (p mm) eat() string {
	return "pk eat"
}
func Test25(t *testing.T) {
	animals := []animaler{pkq{}, mm{}}
	for i := 0; i < 12; i++ {
		num := rand.Intn(2)
		num2 := rand.Intn(2)
		animal := animals[num]
		str := ""
		switch num2 {
		case 0:
			str = animal.eat()
		case 1:
			str = animal.talk()
		}
		fmt.Println(str)
	}
}
