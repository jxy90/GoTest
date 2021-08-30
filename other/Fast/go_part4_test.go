package Fast

import (
	"fmt"
	"testing"
)

func transform(planets [2]string) {
	for i := range planets {
		planets[i] = "New" + planets[i]
	}
}
func transform2(planets []string) {
	for i := range planets {
		planets[i] = "New" + planets[i]
	}
}

func Test16(t *testing.T) {
	fmt.Println("----------")
	planets := [...]string{"123", "123"}
	fmt.Println(planets)
	fmt.Println("----------")
	planets1 := [...]string{"123", "123"}
	transform(planets1)
	fmt.Println(planets1)
	planets2 := make([]string, 0)
	planets2 = append(planets2, "123")
	planets2 = append(planets2, "234")
	transform2(planets2)
	fmt.Println(planets2)
	fmt.Println("----------")
	dump := func(label string, slice []string) {
		fmt.Printf("%v : length %v,capacity %v %v \n", label, len(slice), cap(slice), slice)
	}
	dwarfs := []string{"1", "2", "3", "4", "5"}
	dump("dwarfa", dwarfs)
	dump("dwarfa[1:2]", dwarfs[1:2])
	dump("dwarfa[3:4]", dwarfs[3:4])
	fmt.Println("----------")
	//底层数组不同
	planets3 := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	terrestrial := planets3[0:4:4]
	worlds := append(terrestrial, "2")
	dump("planets3", planets3)
	dump("terrestrial", terrestrial)
	dump("worlds", worlds)
}

func Test19(t *testing.T) {
	fmt.Println("----------")
	temerature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	fmt.Println(temerature)
	fmt.Println("----------")
}
