package main

import (
	"encoding/json"
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
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	sl1 := arr[:5]
	fmt.Println(len(sl1), cap(sl1))
	sl1 = append(sl1, 7)
	fmt.Println(arr)
	sl2 := arr[:5:5]
	fmt.Println(len(sl2), cap(sl2))
	//p1 := People{"jxy1", 1, nil}
	//p2 := &People{"jxy2", 2, nil}
	//p1.getName()
	//p1.getAge()
	//p2.getName()
	//p2.getAge()
	//s1 := make([]int, 10)
	//m1 := map[string]int{}
	//fmt.Println(s1)
	//fmt.Println(m1)
	//fmt.Println(p1)
	//change(s1, m1, &p1)
	//fmt.Println(s1)
	//fmt.Println(m1)
	//fmt.Println(p1)

	item := &GuardRuleInput{
		AlerterType:      1,
		TargetMax:        0,
		TargetMin:        0,
		CoolDownDuration: 0,
		Duration:         0,
		Interval:         0,
		AreaList:         nil,
		Limit:            0,
		Enabled:          0,
	}

	fmt.Println(item)
	res, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	fmt.Println(string(res))

}

func change(s1 []int, m1 map[string]int, people *People) {
	s1[0] = 1
	m1["1"] = 1
	people.age = 2
	people.hobby = []string{"1"}
}
