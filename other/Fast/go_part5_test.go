package Fast

import (
	"encoding/json"
	"fmt"
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

func Test22(t *testing.T) {
	fmt.Println("----------")
	fmt.Println("----------")
	fmt.Println("----------")
}
